import re

# Define a mapping from percentage score to GPA score based on the typical conversion table
def percentage_to_gpa(percentage):
    if 90 <= percentage <= 100:
        return 4.0
    elif 85 <= percentage <= 89:
        return 3.7
    elif 82 <= percentage <= 84:
        return 3.3
    elif 78 <= percentage <= 81:
        return 3.0
    elif 75 <= percentage <= 77:
        return 2.7
    elif 72 <= percentage <= 74:
        return 2.3
    elif 68 <= percentage <= 71:
        return 2.0
    elif 64 <= percentage <= 67:
        return 1.7
    elif 60 <= percentage <= 63:
        return 1.3
    else:
        return 0.0


def stevens_to_gpa(percentage):
    if 90 <= percentage <= 100:
        return 4.0
    elif 80 <= percentage <= 89:
        return 3.0
    elif 70 <= percentage <= 79:
        return 2.0
    elif 60 <= percentage <= 69:
        return 1.0
    else:
        return 0.0


def wes2_to_gpa(percentage):
    if 85 <= percentage <= 100:
        return 4.0
    elif 70 <= percentage <= 84:
        return 3.0
    elif 60 <= percentage <= 69:
        return 2.0
    else:
        return 0.0

def wes3_to_gpa(percentage):
    if 85 <= percentage <= 100:
        return 4.0
    elif 75 <= percentage <= 84:
        return 3.0
    elif 60 < percentage <= 74:
        return 2.0
    else:
        return 0.0


# Define a mapping from percentage score to GPA score based on the typical conversion table
def percentage_to_gpa_other(percentage):
    if 94 <= percentage <= 100:
        return 4.0
    elif 90 <= percentage <= 93:
        return 3.7
    elif 87 <= percentage <= 89:
        return 3.3
    elif 84 <= percentage <= 86:
        return 3.0
    elif 80 <= percentage <= 83:
        return 2.7
    elif 77 <= percentage <= 79:
        return 2.3
    elif 74 <= percentage <= 76:
        return 2.0
    elif 70 <= percentage <= 73:
        return 1.7
    elif 67 <= percentage <= 69:
        return 1.3
    elif 64 <= percentage <= 66:
        return 1.0
    elif 60 <= percentage <= 63:
        return 0.7
    else:
        return 0.0


        
def wes_to_gpa(percentage):
    if 95 <= percentage <= 100:
        return 4.0
    elif 75 <= percentage <= 94:
        return 3.0
    elif 60 < percentage <= 74:
        return 2.0
    elif percentage <= 60:
        return 1.0
    else:
        return 0.0  


# Manually transcribed data from the image based on observed patterns
# Format: [(score, credit), ...]
data = [
    (76, 4), (91, 6), (73, 2), (93, 4), (94, 1), (80, 3), (97, 4), (85, 3), (84, 2), 
    (96, 3), (86, 4), (92, 1), (83, 5), (82, 4), (93, 2), (88, 4), (64, 3), (79, 3), 
    (72, 4), (87, 4), (87, 2), (92, 2), (85, 4), (95, 1), (82, 3), (60, 4), (88, 3), 
    (64, 4), (89, 4), (93, 1), (90, 4), (60, 4), (73, 4), (86, 3), (89, 4), (99, 2), 
    (65, 4), (82, 8), (85, 6), (96, 6), (95, 2), (85, 3), (83, 2), (85, 4), (90, 6), 
    (88, 1), (83, 3), (89, 2), (88, 3), (79, 4), (93, 4), (88, 2), (69, 4), (79, 2), 
    (77, 5), (72, 3), (89, 2), (86, 6), (86, 4), (91, 4), (74, 2), (96, 4), (70, 4), 
    (82, 6)
]
print('data len ',len(data))

# Calculate weighted GPA
total_weighted_gpa = 0
total_credits = 0

for score, credit in data:
    gpa = wes3_to_gpa(score)
    total_weighted_gpa += gpa * credit
    total_credits += credit


print('total_weighted_gpa  ',total_weighted_gpa)
print('total_credits  ',total_credits)
# Final GPA calculation
final_gpa = total_weighted_gpa / total_credits if total_credits > 0 else 0
final_gpa

print('final_gpa  ',final_gpa)



def calculate_cgpa(data):
    """
    计算 CGPA (Cumulative Grade Point Average)
    
    参数:
    - data: 列表，每个元素为一个元组，(成绩, 学分)
    
    返回:
    - CGPA，保留两位小数
    """
    total_points = 0
    total_credits = 0

    for score, credit in data:
        total_points += score * credit
        total_credits += credit

    # 确保总学分不为 0
    if total_credits == 0:
        raise ValueError("总学分不能为 0")
    
    cgpa = total_points / total_credits
    return round(cgpa, 2)

# 计算 CGPA
try:
    result = calculate_cgpa(data)
    print(f"CGPA: {result}")
except ValueError as e:
    print(e)