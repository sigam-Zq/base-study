package core

import (
	"encoding/json"
	"fmt"
	"os"
	"plugin-sche/schema"
	"sync"
)

// StatusManager defines the interface for managing plugin states.
type StatusManager interface {
	Get(name string) (*schema.PluginState, bool)
	Set(name string, state *schema.PluginState)
	GetAll() map[string]*schema.PluginState
	Save() error
	Load() error
	UpdatePluginState(name string, updateFunc func(state *schema.PluginState)) error
	Reset() error
}

// FileStatusManager implements StatusManager using a JSON file.
type FileStatusManager struct {
	filePath string
	states   map[string]*schema.PluginState
	mu       sync.RWMutex
}

// NewFileStatusManager creates a new FileStatusManager and loads initial states from the file.
func NewFileStatusManager(filePath string) (*FileStatusManager, error) {
	fsm := &FileStatusManager{
		filePath: filePath,
		states:   make(map[string]*schema.PluginState),
	}
	if err := fsm.Load(); err != nil {
		return nil, err
	}
	return fsm, nil
}

// Get retrieves the state of a specific plugin.
func (fsm *FileStatusManager) Get(name string) (*schema.PluginState, bool) {
	fsm.mu.RLock()
	defer fsm.mu.RUnlock()
	state, ok := fsm.states[name]
	return state, ok
}

// Set adds or updates the state of a specific plugin.
func (fsm *FileStatusManager) Set(name string, state *schema.PluginState) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	fsm.states[name] = state
}

// GetAll retrieves all plugin states.
func (fsm *FileStatusManager) GetAll() map[string]*schema.PluginState {
	fsm.mu.RLock()
	defer fsm.mu.RUnlock()
	// Return a copy to prevent external modification
	statesCopy := make(map[string]*schema.PluginState)
	for name, state := range fsm.states {
		statesCopy[name] = state
	}
	return statesCopy
}

// Save persists the current plugin states to the JSON file.
func (fsm *FileStatusManager) Save() error {
	data, err := json.MarshalIndent(fsm.states, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fsm.filePath, data, 0644)
}

// Load reads the plugin states from the JSON file.
func (fsm *FileStatusManager) Load() error {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()

	data, err := os.ReadFile(fsm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, initialize with an empty map
			fsm.states = make(map[string]*schema.PluginState)
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &fsm.states)
}

// UpdatePluginState updates the state of a specific plugin.
func (fsm *FileStatusManager) UpdatePluginState(name string, updateFunc func(state *schema.PluginState)) error {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()

	state, ok := fsm.states[name]
	if !ok {
		// If the plugin state doesn't exist, create a default one
		state = &schema.PluginState{
			// You might want to set some default values here
			Status: "idle",
			// Enabled: true, // Default to enabled
		}
		fsm.states[name] = state
	}

	updateFunc(state)

	return fsm.Save()
}

// Reset clears all plugin states and saves an empty state file.
func (fsm *FileStatusManager) Reset() error {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()

	fsm.states = make(map[string]*schema.PluginState) // Clear in-memory states

	if _, err := os.Stat(fsm.filePath); err == os.ErrNotExist {
		os.Mkdir(fsm.filePath, 0o766)
	}

	// Attempt to remove the file, if it exists. If not, it's fine.
	if err := os.Remove(fsm.filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("无法删除状态文件 %s: %w", fsm.filePath, err)
	}

	// Save an empty state file (or let Load handle it if it doesn't exist)
	return fsm.Save()
}
