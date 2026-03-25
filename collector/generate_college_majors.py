#!/usr/bin/env python3
"""
生成高校-专业关联数据
基于1300所高校和830个专业
"""
import json
import os
import random

random.seed(42)

def save_json(data, filename):
    filepath = os.path.join('../data', filename)
    os.makedirs(os.path.dirname(filepath), exist_ok=True)
    with open(filepath, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f'Saved: {filepath}')
    return filepath

# 加载数据
with open('../data/colleges.json', 'r', encoding='utf-8') as f:
    colleges = json.load(f)

with open('../data/majors_845.json', 'r', encoding='utf-8') as f:
    majors = json.load(f)

# 专业适配类型
MAJOR_SUITABILITY = {
    "哲学": ["Comprehensive"],
    "经济学": ["Comprehensive", "Finance & Economics"],
    "法学": ["Comprehensive", "Politics & Law"],
    "教育学": ["Normal"],
    "文学": ["Comprehensive", "Normal", "Language"],
    "历史学": ["Comprehensive", "Normal"],
    "理学": ["Comprehensive", "Normal", "Science & Engineering"],
    "工学": ["Science & Engineering"],
    "农学": ["Agriculture"],
    "医学": ["Medicine", "Comprehensive"],
    "管理学": ["Comprehensive", "Finance & Economics"],
    "艺术学": ["Arts", "Comprehensive"],
}

def get_suitable_colleges(major):
    discipline = major["discipline"]
    suitable_types = MAJOR_SUITABILITY.get(discipline, ["Comprehensive", "Science & Engineering"])
    
    # 根据专业类别筛选高校
    filtered = [c for c in colleges if c["type"] in suitable_types]
    if not filtered:
        filtered = colleges
    return filtered

def generate_college_major_associations():
    associations = []
    assoc_id = 1
    
    print(f"Generating college-major associations for {len(colleges)} colleges and {len(majors)} majors...")
    
    for college in colleges:
        college_rank = college["rankings"]["comprehensive"]
        college_type = college["type"]
        
        # 根据高校排名确定开设专业数量
        # 排名越前，开设专业越多
        if college_rank <= 50:
            num_majors = random.randint(60, 90)
        elif college_rank <= 100:
            num_majors = random.randint(40, 70)
        elif college_rank <= 200:
            num_majors = random.randint(30, 50)
        elif college_rank <= 500:
            num_majors = random.randint(20, 40)
        else:
            num_majors = random.randint(10, 30)
        
        # 筛选适合的专业
        suitable_majors = [m for m in majors if college_type in MAJOR_SUITABILITY.get(m["discipline"], ["Comprehensive"])]
        if not suitable_majors:
            suitable_majors = majors
        
        # 优先选择同类专业
        priority_majors = [m for m in suitable_majors if m["discipline"] == college_type]
        if len(priority_majors) < num_majors // 2:
            priority_majors = suitable_majors[:num_majors]
        
        # 随机选择专业
        selected = random.sample(suitable_majors, min(num_majors, len(suitable_majors)))
        
        for major in selected:
            # 生成专业信息
            assoc = {
                "id": f"{assoc_id:07d}",
                "college_id": college["id"],
                "college_name": college["name"],
                "province": college["province"],
                "city": college["city"],
                "major_code": major["code"],
                "major_name": major["name"],
                "discipline": major["discipline"],
                "category": major["category"],
                "degree": major["degree"],
                "duration": major["duration"],
                "college_rank": college_rank,
            }
            associations.append(assoc)
            assoc_id += 1
        
        if assoc_id % 10000 == 0:
            print(f"  Processed {assoc_id} associations...")
    
    return associations

def main():
    print("="*60)
    print("Generating College-Major Associations")
    print("="*60)
    
    associations = generate_college_major_associations()
    print(f"\nTotal associations: {len(associations)}")
    
    # 按高校统计
    college_counts = {}
    for a in associations:
        cid = a["college_id"]
        college_counts[cid] = college_counts.get(cid, 0) + 1
    
    print(f"Colleges with majors: {len(college_counts)}")
    print(f"Average majors per college: {len(associations) // len(college_counts)}")
    
    # 保存数据
    save_json(associations, 'college_majors_full.json')
    
    # 生成汇总统计
    stats = {
        "total_colleges": len(colleges),
        "total_majors": len(majors),
        "total_associations": len(associations),
        "avg_majors_per_college": len(associations) // len(college_counts),
    }
    save_json(stats, 'college_majors_stats.json')
    
    print("\nSample associations:")
    for a in associations[:3]:
        print(f"  {a['college_name']} - {a['major_name']}")

if __name__ == '__main__':
    main()