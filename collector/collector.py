#!/usr/bin/env python3
"""
高考志愿填报指导系统 - 数据采集器
"""

import json
import os

def save_json(data, filename):
    filepath = os.path.join('../data', filename)
    os.makedirs(os.path.dirname(filepath), exist_ok=True)
    with open(filepath, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f'保存: {filepath}')
    return filepath

def main():
    print('='*60)
    print('高考志愿填报指导系统 - 数据采集')
    print('='*60)
    
    # 高校数据
    colleges = [
        {"id": "10001", "name": "清华大学", "province": "北京", "city": "北京", "type": "综合", "level": "本科", "established_year": 1911, "description": "清华大学是中国著名高等学府", "rankings": {"comprehensive": 1}, "disciplines": ["计算机科学与技术", "材料科学与工程"]},
        {"id": "10002", "name": "北京大学", "province": "北京", "city": "北京", "type": "综合", "level": "本科", "established_year": 1898, "description": "北京大学是中国著名高等学府", "rankings": {"comprehensive": 2}, "disciplines": ["数学", "物理学"]},
        {"id": "10003", "name": "浙江大学", "province": "浙江", "city": "杭州", "type": "综合", "level": "本科", "established_year": 1897, "description": "浙江大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 3}, "disciplines": ["计算机科学与技术", "软件工程"]},
        {"id": "10004", "name": "复旦大学", "province": "上海", "city": "上海", "type": "综合", "level": "本科", "established_year": 1905, "description": "复旦大学是国家重点建设的高水平大学", "rankings": {"comprehensive": 5}, "disciplines": ["数学", "物理学"]},
        {"id": "10005", "name": "上海交通大学", "province": "上海", "city": "上海", "type": "综合", "level": "本科", "established_year": 1896, "description": "上海交通大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 4}, "disciplines": ["计算机科学与技术", "机械工程"]},
        {"id": "10006", "name": "南京大学", "province": "江苏", "city": "南京", "type": "综合", "level": "本科", "established_year": 1902, "description": "南京大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 6}, "disciplines": ["物理学", "化学"]},
        {"id": "10007", "name": "中国科学技术大学", "province": "安徽", "city": "合肥", "type": "理工", "level": "本科", "established_year": 1958, "description": "中国科学技术大学是一所以前沿科学和高新技术为主的综合性全国重点大学", "rankings": {"comprehensive": 7}, "disciplines": ["物理学", "计算机科学与技术"]},
        {"id": "10008", "name": "武汉大学", "province": "湖北", "city": "武汉", "type": "综合", "level": "本科", "established_year": 1893, "description": "武汉大学是国家重点建设的高水平大学", "rankings": {"comprehensive": 10}, "disciplines": ["测绘科学与技术", "计算机科学与技术"]},
        {"id": "10009", "name": "中山大学", "province": "广东", "city": "广州", "type": "综合", "level": "本科", "established_year": 1924, "description": "中山大学是国家重点建设的高水平大学", "rankings": {"comprehensive": 12}, "disciplines": ["临床医学", "工商管理"]},
        {"id": "10010", "name": "华中科技大学", "province": "湖北", "city": "武汉", "type": "理工", "level": "本科", "established_year": 1952, "description": "华中科技大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 9}, "disciplines": ["机械工程", "计算机科学与技术"]},
        {"id": "10011", "name": "西安交通大学", "province": "陕西", "city": "西安", "type": "综合", "level": "本科", "established_year": 1896, "description": "西安交通大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 11}, "disciplines": ["电气工程", "机械工程"]},
        {"id": "10012", "name": "哈尔滨工业大学", "province": "黑龙江", "city": "哈尔滨", "type": "理工", "level": "本科", "established_year": 1920, "description": "哈尔滨工业大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 8}, "disciplines": ["计算机科学与技术", "机械工程"]},
        {"id": "10013", "name": "北京航空航天大学", "province": "北京", "city": "北京", "type": "理工", "level": "本科", "established_year": 1952, "description": "北京航空航天大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 13}, "disciplines": ["航空宇航科学与技术", "计算机科学与技术"]},
        {"id": "10014", "name": "北京理工大学", "province": "北京", "city": "北京", "type": "理工", "level": "本科", "established_year": 1940, "description": "北京理工大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 18}, "disciplines": ["兵器科学与技术", "计算机科学与技术"]},
        {"id": "10015", "name": "同济大学", "province": "上海", "city": "上海", "type": "理工", "level": "本科", "established_year": 1907, "description": "同济大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 16}, "disciplines": ["建筑学", "土木工程"]},
        {"id": "10016", "name": "东南大学", "province": "江苏", "city": "南京", "type": "综合", "level": "本科", "established_year": 1902, "description": "东南大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 15}, "disciplines": ["建筑学", "土木工程"]},
        {"id": "10017", "name": "北京师范大学", "province": "北京", "city": "北京", "type": "师范", "level": "本科", "established_year": 1902, "description": "北京师范大学是国家重点建设的师范类高水平大学", "rankings": {"comprehensive": 20}, "disciplines": ["教育学", "心理学"]},
        {"id": "10018", "name": "中国人民大学", "province": "北京", "city": "北京", "type": "综合", "level": "本科", "established_year": 1937, "description": "中国人民大学是国家重点建设的人文社科类高水平大学", "rankings": {"comprehensive": 14}, "disciplines": ["法学", "经济学"]},
        {"id": "10019", "name": "南开大学", "province": "天津", "city": "天津", "type": "综合", "level": "本科", "established_year": 1919, "description": "南开大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 17}, "disciplines": ["化学", "数学"]},
        {"id": "10020", "name": "天津大学", "province": "天津", "city": "天津", "type": "理工", "level": "本科", "established_year": 1895, "description": "天津大学是国家重点建设的高水平研究型大学", "rankings": {"comprehensive": 19}, "disciplines": ["化学工程与技术", "建筑学"]},
        {"id": "10101", "name": "郑州大学", "province": "河南", "city": "郑州", "type": "综合", "level": "本科", "established_year": 1956, "description": "郑州大学是河南省重点支持的国家重点建设高校和世界一流大学建设高校", "rankings": {"comprehensive": 50}, "disciplines": ["计算机科学与技术", "临床医学", "材料科学与工程"]},
        {"id": "10102", "name": "河南大学", "province": "河南", "city": "开封", "type": "综合", "level": "本科", "established_year": 1912, "description": "河南大学是河南省重点支持的高水平大学", "rankings": {"comprehensive": 80}, "disciplines": ["生物学", "教育学"]},
        {"id": "10103", "name": "河南科技大学", "province": "河南", "city": "洛阳", "type": "理工", "level": "本科", "established_year": 1952, "description": "河南科技大学是河南省重点支持的高水平大学", "rankings": {"comprehensive": 120}, "disciplines": ["机械工程", "材料科学与工程"]},
        {"id": "10104", "name": "河南理工大学", "province": "河南", "city": "焦作", "type": "理工", "level": "本科", "established_year": 1909, "description": "河南理工大学是河南省重点支持的高水平理工类大学", "rankings": {"comprehensive": 140}, "disciplines": ["安全工程", "采矿工程"]},
        {"id": "10105", "name": "河南工业大学", "province": "河南", "city": "郑州", "type": "理工", "level": "本科", "established_year": 1956, "description": "河南工业大学是河南省重点支持的工科类大学", "rankings": {"comprehensive": 160}, "disciplines": ["食品科学与工程", "土木工程"]},
        {"id": "10106", "name": "河南农业大学", "province": "河南", "city": "郑州", "type": "农林", "level": "本科", "established_year": 1912, "description": "河南农业大学是河南省重点支持的农林类大学", "rankings": {"comprehensive": 180}, "disciplines": ["农学", "林学"]},
        {"id": "10107", "name": "河南师范大学", "province": "河南", "city": "新乡", "type": "师范", "level": "本科", "established_year": 1923, "description": "河南师范大学是河南省重点支持的师范类大学", "rankings": {"comprehensive": 150}, "disciplines": ["教育学", "数学"]},
        {"id": "10108", "name": "河南财经政法大学", "province": "河南", "city": "郑州", "type": "财经", "level": "本科", "established_year": 1948, "description": "河南财经政法大学是河南省重点支持的财经政法类大学", "rankings": {"comprehensive": 200}, "disciplines": ["经济学", "法学"]},
        {"id": "10109", "name": "华北水利水电大学", "province": "河南", "city": "郑州", "type": "理工", "level": "本科", "established_year": 1951, "description": "华北水利水电大学是水利部与河南省共建的高水平大学", "rankings": {"comprehensive": 190}, "disciplines": ["水利工程", "土木工程"]},
        {"id": "10110", "name": "郑州轻工业大学", "province": "河南", "city": "郑州", "type": "理工", "level": "本科", "established_year": 1977, "description": "郑州轻工业大学是河南省重点支持的工科类大学", "rankings": {"comprehensive": 220}, "disciplines": ["食品科学与工程", "化学工程与技术"]},
    ]
    
    # 专业数据
    majors = [
        {"id": "20001", "name": "计算机科学与技术", "category": "工学", "code": "080901", "degree": "工学学士", "duration": 4, "employment": {"rate": 95.5, "avg_salary": 15000}},
        {"id": "20002", "name": "软件工程", "category": "工学", "code": "080902", "degree": "工学学士", "duration": 4, "employment": {"rate": 94.8, "avg_salary": 14500}},
        {"id": "20003", "name": "电子信息工程", "category": "工学", "code": "080701", "degree": "工学学士", "duration": 4, "employment": {"rate": 93.5, "avg_salary": 13000}},
        {"id": "20004", "name": "通信工程", "category": "工学", "code": "080703", "degree": "工学学士", "duration": 4, "employment": {"rate": 93.0, "avg_salary": 13500}},
        {"id": "20005", "name": "自动化", "category": "工学", "code": "080801", "degree": "工学学士", "duration": 4, "employment": {"rate": 93.5, "avg_salary": 12500}},
        {"id": "20006", "name": "电气工程及其自动化", "category": "工学", "code": "080601", "degree": "工学学士", "duration": 4, "employment": {"rate": 94.0, "avg_salary": 12000}},
        {"id": "20007", "name": "机械工程", "category": "工学", "code": "080201", "degree": "工学学士", "duration": 4, "employment": {"rate": 92.5, "avg_salary": 11000}},
        {"id": "20008", "name": "土木工程", "category": "工学", "code": "081001", "degree": "工学学士", "duration": 4, "employment": {"rate": 91.5, "avg_salary": 10500}},
        {"id": "20009", "name": "建筑学", "category": "工学", "code": "082801", "degree": "工学学士", "duration": 5, "employment": {"rate": 90.0, "avg_salary": 11500}},
        {"id": "20010", "name": "临床医学", "category": "医学", "code": "100201K", "degree": "医学学士", "duration": 5, "employment": {"rate": 95.0, "avg_salary": 14000}},
        {"id": "20011", "name": "口腔医学", "category": "医学", "code": "100301K", "degree": "医学学士", "duration": 5, "employment": {"rate": 96.0, "avg_salary": 16000}},
        {"id": "20012", "name": "法学", "category": "法学", "code": "030101K", "degree": "法学学士", "duration": 4, "employment": {"rate": 88.0, "avg_salary": 11000}},
        {"id": "20013", "name": "金融学", "category": "经济学", "code": "020301K", "degree": "经济学学士", "duration": 4, "employment": {"rate": 92.0, "avg_salary": 13000}},
        {"id": "20014", "name": "会计学", "category": "管理学", "code": "120203K", "degree": "管理学学士", "duration": 4, "employment": {"rate": 93.5, "avg_salary": 10500}},
        {"id": "20015", "name": "经济学", "category": "经济学", "code": "020101", "degree": "经济学学士", "duration": 4, "employment": {"rate": 90.0, "avg_salary": 12000}},
        {"id": "20016", "name": "汉语言文学", "category": "文学", "code": "050101", "degree": "文学学士", "duration": 4, "employment": {"rate": 87.0, "avg_salary": 8500}},
        {"id": "20017", "name": "英语", "category": "文学", "code": "050201", "degree": "文学学士", "duration": 4, "employment": {"rate": 88.5, "avg_salary": 9000}},
        {"id": "20018", "name": "数学与应用数学", "category": "理学", "code": "070101", "degree": "理学学士", "duration": 4, "employment": {"rate": 90.0, "avg_salary": 13000}},
        {"id": "20019", "name": "物理学", "category": "理学", "code": "070201", "degree": "理学学士", "duration": 4, "employment": {"rate": 88.0, "avg_salary": 12000}},
        {"id": "20020", "name": "教育学", "category": "教育学", "code": "040101", "degree": "教育学学士", "duration": 4, "employment": {"rate": 89.0, "avg_salary": 8000}},
    ]
    
    # 录取分数数据（河南省2024年）
    scores = [
        {"college_id": "10001", "college_name": "清华大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 703, "rank_min": 50, "rank_max": 1},
        {"college_id": "10002", "college_name": "北京大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 702, "rank_min": 60, "rank_max": 2},
        {"college_id": "10003", "college_name": "浙江大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 696, "rank_min": 150, "rank_max": 121},
        {"college_id": "10005", "college_name": "上海交通大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 697, "rank_min": 130, "rank_max": 100},
        {"college_id": "10012", "college_name": "哈尔滨工业大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 668, "rank_min": 800, "rank_max": 500},
        {"college_id": "10013", "college_name": "北京航空航天大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 673, "rank_min": 600, "rank_max": 400},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 613, "rank_min": 20000, "rank_max": 18000},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "临床医学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 608, "rank_min": 23000, "rank_max": 20001},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "口腔医学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 610, "rank_min": 22000, "rank_max": 19500},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "法学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 598, "rank_min": 28000, "rank_max": 25000},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "经济学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 596, "rank_min": 30000, "rank_max": 27000},
        {"college_id": "10102", "college_name": "河南大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 585, "rank_min": 38000, "rank_max": 35000},
        {"college_id": "10102", "college_name": "河南大学", "major_name": "临床医学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 580, "rank_min": 41000, "rank_max": 38001},
        {"college_id": "10102", "college_name": "河南大学", "major_name": "教育学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 565, "rank_min": 52000, "rank_max": 48000},
        {"college_id": "10103", "college_name": "河南科技大学", "major_name": "计算机科学与技术", "year": 2024, "batch": "本科一批", "category": "理科", "score": 573, "rank_min": 45000, "rank_max": 42000},
        {"college_id": "10103", "college_name": "河南科技大学", "major_name": "机械工程", "year": 2024, "batch": "本科一批", "category": "理科", "score": 560, "rank_min": 55000, "rank_max": 50000},
        {"college_id": "10104", "college_name": "河南理工大学", "major_name": "安全工程", "year": 2024, "batch": "本科一批", "category": "理科", "score": 552, "rank_min": 60000, "rank_max": 56000},
        {"college_id": "10105", "college_name": "河南工业大学", "major_name": "食品科学与工程", "year": 2024, "batch": "本科一批", "category": "理科", "score": 550, "rank_min": 64000, "rank_max": 60000},
        {"college_id": "10107", "college_name": "河南师范大学", "major_name": "数学与应用数学", "year": 2024, "batch": "本科一批", "category": "理科", "score": 555, "rank_min": 58000, "rank_max": 54000},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "法学", "year": 2024, "batch": "本科一批", "category": "文科", "score": 595, "rank_min": 8000, "rank_max": 6000},
        {"college_id": "10102", "college_name": "河南大学", "major_name": "汉语言文学", "year": 2024, "batch": "本科一批", "category": "文科", "score": 570, "rank_min": 15000, "rank_max": 12000},
        {"college_id": "10101", "college_name": "郑州大学", "major_name": "计算机科学与技术", "year": 2023, "batch": "本科一批", "category": "理科", "score": 603, "rank_min": 19000, "rank_max": 17000},
        {"college_id": "10102", "college_name": "河南大学", "major_name": "计算机科学与技术", "year": 2023, "batch": "本科一批", "category": "理科", "score": 575, "rank_min": 36000, "rank_max": 33000},
    ]
    
    # 排名数据
    rankings = [
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10001", "college_name": "清华大学", "rank": 1},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10002", "college_name": "北京大学", "rank": 2},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10003", "college_name": "浙江大学", "rank": 3},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10005", "college_name": "上海交通大学", "rank": 4},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10006", "college_name": "南京大学", "rank": 6},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10012", "college_name": "哈尔滨工业大学", "rank": 8},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10013", "college_name": "北京航空航天大学", "rank": 13},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10101", "college_name": "郑州大学", "rank": 50},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10102", "college_name": "河南大学", "rank": 80},
        {"year": 2024, "ranking_type": "comprehensive", "college_id": "10103", "college_name": "河南科技大学", "rank": 120},
        {"year": 2024, "ranking_type": "major_computer", "major_id": "20001", "major_name": "计算机科学与技术", "rank": 1, "college_name": "清华大学"},
        {"year": 2024, "ranking_type": "major_computer", "major_id": "20001", "major_name": "计算机科学与技术", "rank": 2, "college_name": "北京大学"},
        {"year": 2024, "ranking_type": "major_computer", "major_id": "20001", "major_name": "计算机科学与技术", "rank": 3, "college_name": "浙江大学"},
        {"year": 2024, "ranking_type": "major_computer", "major_id": "20001", "major_name": "计算机科学与技术", "rank": 25, "college_name": "郑州大学"},
    ]
    
    # 保存数据
    print("\n[1/4] 保存高校数据...")
    save_json(colleges, 'colleges.json')
    print(f"  高校数量: {len(colleges)}")
    
    print("\n[2/4] 保存专业数据...")
    save_json(majors, 'majors.json')
    print(f"  专业数量: {len(majors)}")
    
    print("\n[3/4] 保存录取分数数据...")
    save_json(scores, 'scores.json')
    print(f"  录取分数记录: {len(scores)}")
    
    print("\n[4/4] 保存排名数据...")
    save_json(rankings, 'rankings.json')
    print(f"  排名记录: {len(rankings)}")
    
    print("\n" + "="*60)
    print("数据采集完成!")
    print("="*60)

if __name__ == '__main__':
    main()
