#!/usr/bin/env python3
import json
import os
import random

def save_json(data, filename):
    filepath = os.path.join('../data', filename)
    os.makedirs(os.path.dirname(filepath), exist_ok=True)
    with open(filepath, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f'Saved: {filepath}')
    return filepath

# Major database with categories and suitable college types
MAJORS_DB = {
    # Engineering majors
    "Computer Science and Technology": {"category": "Engineering", "code": "080901", "types": ["Comprehensive", "Science & Engineering", "Normal"]},
    "Software Engineering": {"category": "Engineering", "code": "080902", "types": ["Comprehensive", "Science & Engineering"]},
    "Electronic Information Engineering": {"category": "Engineering", "code": "080701", "types": ["Science & Engineering", "Comprehensive"]},
    "Communication Engineering": {"category": "Engineering", "code": "080703", "types": ["Science & Engineering"]},
    "Automation": {"category": "Engineering", "code": "080801", "types": ["Science & Engineering"]},
    "Electrical Engineering and Automation": {"category": "Engineering", "code": "080601", "types": ["Science & Engineering"]},
    "Mechanical Engineering": {"category": "Engineering", "code": "080201", "types": ["Science & Engineering"]},
    "Materials Science and Engineering": {"category": "Engineering", "code": "080401", "types": ["Science & Engineering", "Comprehensive"]},
    "Chemical Engineering": {"category": "Engineering", "code": "081301", "types": ["Science & Engineering"]},
    "Civil Engineering": {"category": "Engineering", "code": "081001", "types": ["Science & Engineering", "Comprehensive"]},
    "Architecture": {"category": "Engineering", "code": "082801", "types": ["Science & Engineering"], "duration": 5},
    "Environmental Engineering": {"category": "Engineering", "code": "082502", "types": ["Science & Engineering"]},
    "Biomedical Engineering": {"category": "Engineering", "code": "082601", "types": ["Science & Engineering", "Medicine"]},
    "Food Science and Engineering": {"category": "Engineering", "code": "082701", "types": ["Agriculture", "Science & Engineering"]},
    "Transportation Engineering": {"category": "Engineering", "code": "081801", "types": ["Science & Engineering"]},
    "Hydraulic Engineering": {"category": "Engineering", "code": "081101", "types": ["Science & Engineering"]},
    "Mining Engineering": {"category": "Engineering", "code": "081501", "types": ["Science & Engineering"]},
    "Safety Engineering": {"category": "Engineering", "code": "082901", "types": ["Science & Engineering"]},
    "Aerospace Engineering": {"category": "Engineering", "code": "082001", "types": ["Science & Engineering"]},
    "Naval Architecture and Ocean Engineering": {"category": "Engineering", "code": "082400", "types": ["Science & Engineering"]},
    "Nuclear Engineering": {"category": "Engineering", "code": "082201", "types": ["Science & Engineering"]},
    "Photonics Information Science and Engineering": {"category": "Engineering", "code": "080307", "types": ["Science & Engineering"]},
    "Instrument Science and Engineering": {"category": "Engineering", "code": "080300", "types": ["Science & Engineering"]},
    "Vehicle Engineering": {"category": "Engineering", "code": "080207", "types": ["Science & Engineering"]},
    "Railway Engineering": {"category": "Engineering", "code": "081002", "types": ["Science & Engineering"]},
    "纺织工程": {"category": "Engineering", "code": "081601", "types": ["Science & Engineering"]},
    "印刷工程": {"category": "Engineering", "code": "081701", "types": ["Science & Engineering"]},
    "包装工程": {"category": "Engineering", "code": "081702", "types": ["Science & Engineering"]},
    "船舶与海洋工程": {"category": "Engineering", "code": "082400", "types": ["Science & Engineering"]},
    "海洋工程与技术": {"category": "Engineering", "code": "082900", "types": ["Science & Engineering"]},
    
    # Science majors
    "Mathematics and Applied Mathematics": {"category": "Science", "code": "070101", "types": ["Comprehensive", "Normal", "Science & Engineering"]},
    "Physics": {"category": "Science", "code": "070201", "types": ["Comprehensive", "Normal", "Science & Engineering"]},
    "Chemistry": {"category": "Science", "code": "070201", "types": ["Comprehensive", "Normal", "Science & Engineering"]},
    "Applied Chemistry": {"category": "Science", "code": "070302", "types": ["Science & Engineering"]},
    "Biology Science": {"category": "Science", "code": "071001", "types": ["Comprehensive", "Agriculture", "Medicine"]},
    "Biotechnology": {"category": "Science", "code": "071002", "types": ["Science & Engineering", "Medicine"]},
    "Statistics": {"category": "Science", "code": "071201", "types": ["Comprehensive", "Science & Engineering"]},
    "Applied Physics": {"category": "Science", "code": "070202", "types": ["Science & Engineering"]},
    "Marine Science": {"category": "Science", "code": "070700", "types": ["Science & Engineering"]},
    "Atmospheric Science": {"category": "Science", "code": "070600", "types": ["Science & Engineering"]},
    "Geology": {"category": "Science", "code": "070901", "types": ["Science & Engineering"]},
    "Geophysics": {"category": "Science", "code": "070801", "types": ["Science & Engineering"]},
    "Psychology": {"category": "Science", "code": "071101", "types": ["Comprehensive", "Normal", "Medicine"]},
    "History": {"category": "Literature", "code": "060101", "types": ["Comprehensive", "Normal"]},
    
    # Medicine majors
    "Clinical Medicine": {"category": "Medicine", "code": "100201K", "types": ["Medicine", "Comprehensive"], "duration": 5},
    "Stomatology": {"category": "Medicine", "code": "100301K", "types": ["Medicine"], "duration": 5},
    "Pharmacy": {"category": "Medicine", "code": "100701", "types": ["Medicine", "Science & Engineering"]},
    "Traditional Chinese Medicine": {"category": "Medicine", "code": "100501K", "types": ["Medicine"], "duration": 5},
    "Nursing": {"category": "Medicine", "code": "101101", "types": ["Medicine"], "duration": 4},
    "Medical Laboratory Technology": {"category": "Medicine", "code": "101001", "types": ["Medicine"]},
    "Public Health and Preventive Medicine": {"category": "Medicine", "code": "100401K", "types": ["Medicine"]},
    "Optical Information Science and Technology": {"category": "Science", "code": "071300", "types": ["Science & Engineering"]},
    "Acupuncture and Moxibustion": {"category": "Medicine", "code": "100502K", "types": ["Medicine"]},
    "Chinese Pharmacology": {"category": "Medicine", "code": "100803", "types": ["Medicine"]},
    
    # Liberal Arts majors
    "Chinese Language and Literature": {"category": "Literature", "code": "050101", "types": ["Comprehensive", "Normal", "Language"]},
    "English": {"category": "Literature", "code": "050201", "types": ["Comprehensive", "Normal", "Language"]},
    "French": {"category": "Literature", "code": "050204", "types": ["Language"]},
    "German": {"category": "Literature", "code": "050203", "types": ["Language"]},
    "Japanese": {"category": "Literature", "code": "050207", "types": ["Language"]},
    "Korean": {"category": "Literature", "code": "050209", "types": ["Language"]},
    "News and Communication": {"category": "Literature", "code": "050301", "types": ["Comprehensive", "Language"]},
    "Broadcasting and Television": {"category": "Literature", "code": "050305", "types": ["Comprehensive", "Arts"]},
    "Law": {"category": "Law", "code": "030101K", "types": ["Comprehensive", "Politics & Law"]},
    "Intellectual Property": {"category": "Law", "code": "030102T", "types": ["Politics & Law"]},
    "Politics and Public Administration": {"category": "Law", "code": "030201", "types": ["Comprehensive", "Politics & Law"]},
    
    # Economics/Management majors
    "Economics": {"category": "Economics", "code": "020101", "types": ["Comprehensive", "Finance & Economics"]},
    "Finance": {"category": "Economics", "code": "020301K", "types": ["Finance & Economics", "Comprehensive"]},
    "International Finance": {"category": "Economics", "code": "020302", "types": ["Finance & Economics"]},
    "Insurance": {"category": "Economics", "code": "020305", "types": ["Finance & Economics"]},
    "Investment": {"category": "Economics", "code": "020304", "types": ["Finance & Economics"]},
    "Accounting": {"category": "Management", "code": "120203K", "types": ["Finance & Economics", "Comprehensive"]},
    "Business Administration": {"category": "Management", "code": "120201K", "types": ["Comprehensive", "Finance & Economics"]},
    "Marketing": {"category": "Management", "code": "120202", "types": ["Finance & Economics"]},
    "Human Resource Management": {"category": "Management", "code": "120206", "types": ["Comprehensive"]},
    "Tourism Management": {"category": "Management", "code": "120901K", "types": ["Comprehensive"]},
    "Hotel Management": {"category": "Management", "code": "120902", "types": ["Comprehensive"]},
    "Information Management and Information Systems": {"category": "Management", "code": "120102", "types": ["Science & Engineering", "Comprehensive"]},
    "E-commerce": {"category": "Management", "code": "120801", "types": ["Comprehensive", "Science & Engineering"]},
    "Logistics Management": {"category": "Management", "code": "120601", "types": ["Comprehensive"]},
    "Public Administration": {"category": "Management", "code": "120401", "types": ["Comprehensive", "Politics & Law"]},
    "Education Management": {"category": "Management", "code": "120402", "types": ["Normal", "Comprehensive"]},
    
    # Education majors
    "Education": {"category": "Education", "code": "040101", "types": ["Normal"], "duration": 4},
    "Primary School Education": {"category": "Education", "code": "040107", "types": ["Normal"]},
    "Preschool Education": {"category": "Education", "code": "040106", "types": ["Normal"]},
    "Physical Education": {"category": "Education", "code": "040201", "types": ["Normal", "Sports"]},
    "Special Education": {"category": "Education", "code": "040108", "types": ["Normal"]},
    "Educational Technology": {"category": "Education", "code": "040104", "types": ["Normal", "Science & Engineering"]},
    
    # Agriculture majors
    "Agronomy": {"category": "Agriculture", "code": "090101", "types": ["Agriculture"]},
    "Plant Protection": {"category": "Agriculture", "code": "090103", "types": ["Agriculture"]},
    "Seed Science and Engineering": {"category": "Agriculture", "code": "090105", "types": ["Agriculture"]},
    "Horticulture": {"category": "Agriculture", "code": "090102", "types": ["Agriculture"]},
    "Forestry": {"category": "Agriculture", "code": "090501", "types": ["Agriculture"]},
    "Animal Science": {"category": "Agriculture", "code": "090301", "types": ["Agriculture"]},
    "Veterinary Medicine": {"category": "Agriculture", "code": "090401", "types": ["Agriculture"]},
    "Grassland Science": {"category": "Agriculture", "code": "090701", "types": ["Agriculture"]},
    "Aquaculture": {"category": "Agriculture", "code": "090601", "types": ["Agriculture"]},
    "Agricultural Mechanization": {"category": "Agriculture", "code": "082302", "types": ["Agriculture", "Science & Engineering"]},
    "Agricultural Economic Management": {"category": "Management", "code": "120301", "types": ["Agriculture", "Comprehensive"]},
    
    # Arts majors
    "Fine Arts": {"category": "Art", "code": "130401", "types": ["Arts"]},
    "Music Performance": {"category": "Art", "code": "130201", "types": ["Arts"]},
    "Musicology": {"category": "Art", "code": "130202", "types": ["Arts"]},
    "Dance Performance": {"category": "Art", "code": "130204", "types": ["Arts"]},
    "Drama and Film": {"category": "Art", "code": "130304", "types": ["Arts", "Comprehensive"]},
    "Design": {"category": "Art", "code": "130500", "types": ["Arts", "Science & Engineering"]},
    "Visual Communication Design": {"category": "Art", "code": "130202", "types": ["Arts"]},
    "Environmental Design": {"category": "Art", "code": "130503", "types": ["Arts"]},
    "Digital Media Art": {"category": "Art", "code": "130508", "types": ["Arts", "Science & Engineering"]},
    
    # Military-related
    "Electronic Countermeasures": {"category": "Military", "code": "110110", "types": ["Military"]},
    "Aerospace Science and Technology": {"category": "Military", "code": "082001", "types": ["Military", "Science & Engineering"]},
    "Military Logistics": {"category": "Military", "code": "110306", "types": ["Military"]},
    "Command and Command Technology": {"category": "Military", "code": "1101", "types": ["Military"]},
}

def get_majors_for_college(college_type, college_rank, city, province):
    selected_majors = []
    
    # Core majors based on college type
    type_majors = {
        "Comprehensive": ["Computer Science and Technology", "Mathematics and Applied Mathematics", "Physics", "Chemistry", "Biology Science", "Chinese Language and Literature", "English", "Economics", "Finance", "Business Administration", "Law", "Education", "History", "Psychology", "Political Science", "Public Administration"],
        "Science & Engineering": ["Computer Science and Technology", "Software Engineering", "Electronic Information Engineering", "Communication Engineering", "Automation", "Electrical Engineering and Automation", "Mechanical Engineering", "Materials Science and Engineering", "Civil Engineering", "Chemical Engineering", "Environmental Engineering", "Mathematics and Applied Mathematics", "Physics", "Biotechnology", "Instrument Science and Engineering"],
        "Normal": ["Chinese Language and Literature", "English", "Mathematics and Applied Mathematics", "Physics", "Chemistry", "Biology Science", "Education", "Primary School Education", "Preschool Education", "Physical Education", "Psychology", "History", "Music Performance", "Fine Arts"],
        "Agriculture": ["Agronomy", "Plant Protection", "Seed Science and Engineering", "Horticulture", "Forestry", "Animal Science", "Veterinary Medicine", "Grassland Science", "Aquaculture", "Agricultural Mechanization", "Food Science and Engineering", "Rural Development"],
        "Medicine": ["Clinical Medicine", "Stomatology", "Pharmacy", "Traditional Chinese Medicine", "Nursing", "Medical Laboratory Technology", "Public Health and Preventive Medicine", "Biomedical Engineering", "Chinese Pharmacology", "Acupuncture and Moxibustion"],
        "Finance & Economics": ["Economics", "Finance", "International Finance", "Insurance", "Investment", "Accounting", "Business Administration", "Marketing", "Human Resource Management", "Tourism Management", "E-commerce", "Logistics Management", "Statistics"],
        "Language": ["Chinese Language and Literature", "English", "French", "German", "Japanese", "Korean", "News and Communication", "Broadcasting and Television", "Education", "Translation"],
        "Arts": ["Fine Arts", "Music Performance", "Musicology", "Dance Performance", "Drama and Film", "Design", "Visual Communication Design", "Environmental Design", "Digital Media Art", "Chinese Language and Literature", "News and Communication"],
        "Sports": ["Physical Education", "Sports Training", "Dance Performance", "Chinese Traditional Sports", "Sports Science"],
        "Politics & Law": ["Law", "Politics and Public Administration", "Intellectual Property", "Public Administration", "Sociology", "Social Work"],
        "Military": ["Electronic Countermeasures", "Aerospace Science and Technology", "Military Logistics", "Command and Command Technology", "Computer Science and Technology", "Electronic Information Engineering"],
    }
    
    base_majors = type_majors.get(college_type, type_majors["Comprehensive"])
    
    # Add more majors for better universities
    num_majors = min(25, 8 + (college_rank // 50) + random.randint(0, 5))
    
    # Select majors
    available = [m for m in base_majors if m in MAJORS_DB]
    selected = random.sample(available, min(num_majors, len(available)))
    
    for major_name in selected:
        major_info = MAJORS_DB[major_name]
        selected_majors.append({
            "name": major_name,
            "category": major_info["category"],
            "code": major_info["code"],
            "degree": f"Bachelor of {major_info['category']}",
            "duration": major_info.get("duration", 4),
            "employment": {
                "rate": round(85 + random.random() * 12, 1),
                "avg_salary": round(6000 + random.random() * 12000 + (college_rank * 20))
            }
        })
    
    return selected_majors

def generate_colleges():
    provinces = [
        ("Beijing", "Beijing", 45),
        ("Shanghai", "Shanghai", 30),
        ("Zhejiang", "Hangzhou", 55),
        ("Jiangsu", "Nanjing", 75),
        ("Guangdong", "Guangzhou", 65),
        ("Shandong", "Jinan", 70),
        ("Henan", "Zhengzhou", 55),
        ("Sichuan", "Chengdu", 55),
        ("Hubei", "Wuhan", 50),
        ("Hunan", "Changsha", 50),
        ("Hebei", "Shijiazhuang", 45),
        ("Anhui", "Hefei", 45),
        ("Fujian", "Fuzhou", 40),
        ("Jiangxi", "Nanchang", 40),
        ("Liaoning", "Shenyang", 45),
        ("Jilin", "Changchun", 35),
        ("Heilongjiang", "Harbin", 35),
        ("Shaanxi", "Xi'an", 45),
        ("Shanxi", "Taiyuan", 35),
        ("Tianjin", "Tianjin", 25),
        ("Chongqing", "Chongqing", 30),
        ("Yunnan", "Kunming", 30),
        ("Guizhou", "Guiyang", 25),
        ("Guangxi", "Nanning", 30),
        ("Hainan", "Haikou", 10),
        ("Inner Mongolia", "Hohhot", 20),
        ("Ningxia", "Yinchuan", 10),
        ("Xinjiang", "Urumqi", 15),
        ("Gansu", "Lanzhou", 20),
        ("Qinghai", "Xining", 5),
        ("Tibet", "Lhasa", 3),
    ]
    
    cities_by_province = {
        "Beijing": ["Beijing"],
        "Shanghai": ["Shanghai"],
        "Zhejiang": ["Hangzhou", "Ningbo", "Wenzhou", "Shaoxing", "Jinhua", "Taizhou", "Zhoushan", "Huzhou", "Jiaxing", "Lishui"],
        "Jiangsu": ["Nanjing", "Suzhou", "Wuxi", "Changzhou", "Yangzhou", "Yancheng", "Xuzhou", "Nantong", "Taizhou", "Lianyungang", "Suqian", "Huai'an", "Zhenjiang"],
        "Guangdong": ["Guangzhou", "Shenzhen", "Foshan", "Dongguan", "Zhongshan", "Zhuhai", "Jiangmen", "Huizhou", "Maoming", "Shaoguan", "Shantou", "Zhanjiang", "Yangjiang", "Qingyuan", "Chaozhou", "Jieyang", "Yunfu", "Meizhou"],
        "Shandong": ["Jinan", "Qingdao", "Yantai", "Weifang", "Zibo", "Jining", "Tai'an", "Linyi", "Dezhou", "Dongying", "Zaozhuang", "Weihai", "Rizhao", "Liaocheng", "Binzhou", "Heze"],
        "Henan": ["Zhengzhou", "Luoyang", "Kaifeng", "Xinxiang", "Puyang", "Anyang", "Hebi", "Jiaozuo", "Sanmenxia", "Nanyang", "Shangqiu", "Xinyang", "Zhoukou", "Zhumadian"],
        "Sichuan": ["Chengdu", "Mianyang", "Deyang", "Nanchong", "Yibin", "Zigong", "Luzhou", "Darl", "Guangyuan", "Suining", "Neijiang", "Leshan", "Nanchong"],
        "Hubei": ["Wuhan", "Yichang", "Xiangyang", "Jingzhou", "Huangshi", "Shiyan", "Xianning", "Xiaogan", "Jingmen", "Ezhou", "Huanggang", "Xiaoogan"],
        "Hunan": ["Changsha", "Zhuzhou", "Xiangtan", "Yueyang", "Changde", "Zhangjiajie", "Yiyang", "Chenzhou", "Huaihua", "Loudi", "Shaoyang", "Yongzhou"],
        "Hebei": ["Shijiazhuang", "Tangshan", "Qinghuangdao", "Handan", "Xingtai", "Baoding", "Zhangjiakou", "Chengde", "Cangzhou", "Langfang", "Hengshui"],
        "Anhui": ["Hefei", "Bengbu", "Wuhu", "Huainan", "Ma'anshan", "Huaibei", "Fuyang", "Anqing", "Huangshan", "Chuzhou", "Liu'an", "Xuancheng", "Chaohu"],
        "Fujian": ["Fuzhou", "Xiamen", "Quanzhou", "Zhangzhou", "Fujian", "Putian", "Sanming", "Longyan", "Ningde"],
        "Jiangxi": ["Nanchang", "Jiujiang", "Jingdezhen", "Pingxiang", "Xinyu", "Yingtan", "Ganzhou", "Yichun", "Shangrao", "Fuzhou"],
        "Liaoning": ["Shenyang", "Dalian", "Anshan", "Fushun", "Liaoyang", "Jinzhou", "Dandong", "Huludao", "Fuxin", "Liaoyang", "Chaoyang"],
        "Jilin": ["Changchun", "Jilin", "Siping", "Liaoyuan", "Tonghua", "Baishan", "Songyuan", "Baicheng", "Yanbian"],
        "Heilongjiang": ["Harbin", "Qiqihar", "Jixi", "Hegang", "Shuangyashan", "Daqing", "Yichun", "Jiamusi", "Heihe", "Suihua", "Mudanjiang"],
        "Shaanxi": ["Xi'an", "Xianyang", "Tongchuan", "Baoji", "Xianyang", "Weinan", "Yan'an", "Hanzhong", "Yulin", "Ankang", "Shangluo"],
        "Shanxi": ["Taiyuan", "Datong", "Yangquan", "Changzhi", "Jincheng", "Shuozhou", "Jinzhong", "Yuncheng", "Linfen", "Lvliang"],
        "Tianjin": ["Tianjin"],
        "Chongqing": ["Chongqing"],
        "Yunnan": ["Kunming", "Dali", "Qujing", "Yuxi", "Honghe", "Wenshan", "Chuxiong", "Zhaotong", "Lijiang", "Pu'er"],
        "Guizhou": ["Guiyang", "Zunyi", "Liupanshui", "Anshun", "Bijie", "Qianxinan"],
        "Guangxi": ["Nanning", "Guilin", "Liuzhou", "Yulin", "Baise", "Hezhou", "Hechi", "Laibin", "Chongzuo", "Fangchenggang", "Beihai"],
        "Hainan": ["Haikou", "Sanya", "Qionghai", "Wenchang", "Sansha"],
        "Inner Mongolia": ["Hohhot", "Baotou", "Wuhai", "Chifeng", "Tongliao", "Ordos", "Hulunbuir", "Bayannur", "Ulanqab", "Xilingol", "Alxa"],
        "Ningxia": ["Yinchuan", "Shizuishan", "Zhongwei", "Guyuan", "Pingluo"],
        "Xinjiang": ["Urumqi", "Shihezi", "Karamay", "Korla", "Aksu", "Hotan", "Turpan", "Changji", "Bortala", "Yili"],
        "Gansu": ["Lanzhou", "Tianshui", "Qingyang", "Pingliang", "Jimentso", "Zhangye", "Wuwei", "Dingxi", "Longnan", "Linxia"],
        "Qinghai": ["Xining", "Haidong", "Haibei", "Hainan", "Golog", "Yushu", "Haixi"],
        "Tibet": ["Lhasa", "Shigatse", "Nyingchi", "Qamdo", "Nagqu"],
    }
    
    types = ["Comprehensive", "Science & Engineering", "Normal", "Agriculture", "Medicine", "Finance & Economics", "Language", "Arts", "Sports", "Politics & Law", "Military"]
    
    colleges = []
    college_id = 1
    
    # Top universities
    top_universities = [
        {"name": "Tsinghua University", "province": "Beijing", "city": "Beijing", "type": "Comprehensive", "rank": 1, "established": 1911},
        {"name": "Peking University", "province": "Beijing", "city": "Beijing", "type": "Comprehensive", "rank": 2, "established": 1898},
        {"name": "Zhejiang University", "province": "Zhejiang", "city": "Hangzhou", "type": "Comprehensive", "rank": 3, "established": 1897},
        {"name": "Shanghai Jiao Tong University", "province": "Shanghai", "city": "Shanghai", "type": "Comprehensive", "rank": 4, "established": 1896},
        {"name": "Fudan University", "province": "Shanghai", "city": "Shanghai", "type": "Comprehensive", "rank": 5, "established": 1905},
        {"name": "Nanjing University", "province": "Jiangsu", "city": "Nanjing", "type": "Comprehensive", "rank": 6, "established": 1902},
        {"name": "University of Science and Technology of China", "province": "Anhui", "city": "Hefei", "type": "Science & Engineering", "rank": 7, "established": 1958},
        {"name": "Harbin Institute of Technology", "province": "Heilongjiang", "city": "Harbin", "type": "Science & Engineering", "rank": 8, "established": 1920},
        {"name": "Huazhong University of Science and Technology", "province": "Hubei", "city": "Wuhan", "type": "Science & Engineering", "rank": 9, "established": 1952},
        {"name": "Wuhan University", "province": "Hubei", "city": "Wuhan", "type": "Comprehensive", "rank": 10, "established": 1893},
        {"name": "Xi'an Jiaotong University", "province": "Shaanxi", "city": "Xi'an", "type": "Comprehensive", "rank": 11, "established": 1896},
        {"name": "Sun Yat-sen University", "province": "Guangdong", "city": "Guangzhou", "type": "Comprehensive", "rank": 12, "established": 1924},
        {"name": "Beihang University", "province": "Beijing", "city": "Beijing", "type": "Science & Engineering", "rank": 13, "established": 1952},
        {"name": "Renmin University of China", "province": "Beijing", "city": "Beijing", "type": "Comprehensive", "rank": 14, "established": 1937},
        {"name": "Southeast University", "province": "Jiangsu", "city": "Nanjing", "type": "Comprehensive", "rank": 15, "established": 1902},
        {"name": "Tongji University", "province": "Shanghai", "city": "Shanghai", "type": "Science & Engineering", "rank": 16, "established": 1907},
        {"name": "Nankai University", "province": "Tianjin", "city": "Tianjin", "type": "Comprehensive", "rank": 17, "established": 1919},
        {"name": "Beijing Institute of Technology", "province": "Beijing", "city": "Beijing", "type": "Science & Engineering", "rank": 18, "established": 1940},
        {"name": "Tianjin University", "province": "Tianjin", "city": "Tianjin", "type": "Science & Engineering", "rank": 19, "established": 1895},
        {"name": "Beijing Normal University", "province": "Beijing", "city": "Beijing", "type": "Normal", "rank": 20, "established": 1902},
        {"name": "Sichuan University", "province": "Sichuan", "city": "Chengdu", "type": "Comprehensive", "rank": 21, "established": 1896},
        {"name": "Shandong University", "province": "Shandong", "city": "Jinan", "type": "Comprehensive", "rank": 22, "established": 1901},
        {"name": "Central South University", "province": "Hunan", "city": "Changsha", "type": "Comprehensive", "rank": 23, "established": 1914},
        {"name": "Xiamen University", "province": "Fujian", "city": "Xiamen", "type": "Comprehensive", "rank": 24, "established": 1921},
        {"name": "Jilin University", "province": "Jilin", "city": "Changchun", "type": "Comprehensive", "rank": 26, "established": 1946},
        {"name": "Northeastern University", "province": "Liaoning", "city": "Shenyang", "type": "Science & Engineering", "rank": 28, "established": 1923},
        {"name": "Dalian University of Technology", "province": "Liaoning", "city": "Dalian", "type": "Science & Engineering", "rank": 27, "established": 1949},
        {"name": "Beijing Jiaotong University", "province": "Beijing", "city": "Beijing", "type": "Science & Engineering", "rank": 30, "established": 1896},
        {"name": "National University of Defense Technology", "province": "Hunan", "city": "Changsha", "type": "Science & Engineering", "rank": 29, "established": 1953},
        {"name": "Northwestern Polytechnical University", "province": "Shaanxi", "city": "Xi'an", "type": "Science & Engineering", "rank": 34, "established": 1938},
        {"name": "Chongqing University", "province": "Chongqing", "city": "Chongqing", "type": "Comprehensive", "rank": 33, "established": 1929},
        {"name": "Lanzhou University", "province": "Gansu", "city": "Lanzhou", "type": "Comprehensive", "rank": 40, "established": 1909},
        {"name": "Suzhou University", "province": "Jiangsu", "city": "Suzhou", "type": "Comprehensive", "rank": 45, "established": 1900},
        {"name": "Ocean University of China", "province": "Shandong", "city": "Qingdao", "type": "Comprehensive", "rank": 48, "established": 1924},
        {"name": "East China Normal University", "province": "Shanghai", "city": "Shanghai", "type": "Normal", "rank": 32, "established": 1951},
        {"name": "Nanjing Normal University", "province": "Jiangsu", "city": "Nanjing", "type": "Normal", "rank": 85, "established": 1902},
        {"name": "Zhengzhou University", "province": "Henan", "city": "Zhengzhou", "type": "Comprehensive", "rank": 50, "established": 1956},
        {"name": "Henan University", "province": "Henan", "city": "Kaifeng", "type": "Comprehensive", "rank": 80, "established": 1912},
        {"name": "Shenzhen University", "province": "Guangdong", "city": "Shenzhen", "type": "Comprehensive", "rank": 68, "established": 1983},
    ]
    
    for uni in top_universities:
        college = {
            "id": f"{college_id:05d}",
            "name": uni["name"],
            "province": uni["province"],
            "city": uni["city"],
            "type": uni["type"],
            "level": "Undergraduate",
            "established_year": uni["established"],
            "departments": 15 + (college_id % 20),
            "faculties": 8000 + (college_id % 20) * 500,
            "website": f"https://www.edu{college_id}.cn",
            "description": f"National key university in {uni['province']}",
            "rankings": {
                "comprehensive": uni["rank"],
                "qs_world": uni["rank"] * 15
            },
            "majors": get_majors_for_college(uni["type"], uni["rank"], uni["city"], uni["province"]),
            "statistics": {
                "enrollment_count": 3000 + (college_id % 50) * 100,
                "employment_rate": 95 + (college_id % 5),
                "average_salary": 10000 + (college_id % 50) * 200
            }
        }
        colleges.append(college)
        college_id += 1
    
    # Generate remaining universities
    for province_name, capital, target_count in provinces:
        cities = cities_by_province.get(province_name, [capital])
        remaining = target_count - len([u for u in top_universities if u["province"] == province_name])
        if remaining <= 0:
            remaining = target_count // 3
        
        for i in range(remaining):
            city = cities[i % len(cities)]
            college_type = types[(college_id + i) % len(types)]
            
            base_rank = len(colleges) + 1
            base_salary = 6000 + (college_id % 40) * 100
            base_enrollment = 2000 + (college_id % 30) * 100
            
            if college_type in ["Medicine", "Finance & Economics"]:
                base_salary += 2000
                base_enrollment -= 300
            elif college_type in ["Agriculture", "Normal"]:
                base_salary -= 1000
                base_enrollment += 500
            
            university_names = [
                f"{capital} University of Technology",
                f"{capital} University of Science",
                f"{capital} Normal University",
                f"{capital} University of Finance",
                f"{capital} Medical University",
                f"{capital} Agricultural University",
                f"{capital} Engineering University",
                f"{capital} College",
                f"{capital} Institute of Technology",
                f"{capital} University of Arts",
                f"{city} University",
                f"{city} Medical College",
                f"{city} Polytechnic",
                f"{city} Vocational University",
            ]
            
            name = university_names[i % len(university_names)]
            if i >= len(university_names):
                name = f"{name} {i // len(university_names) + 1}"
            
            college = {
                "id": f"{college_id:05d}",
                "name": name,
                "province": province_name,
                "city": city,
                "type": college_type,
                "level": "Undergraduate",
                "established_year": 1950 + (college_id % 70),
                "departments": 10 + (college_id % 15),
                "faculties": 4000 + (college_id % 40) * 100,
                "website": f"https://www.{province_name.lower().replace(' ', '')}{college_id}.edu.cn",
                "description": f"University in {city}, {province_name}",
                "rankings": {
                    "comprehensive": base_rank,
                    "qs_world": base_rank * 12
                },
                "majors": get_majors_for_college(college_type, base_rank, city, province_name),
                "statistics": {
                    "enrollment_count": base_enrollment,
                    "employment_rate": 80 + (college_id % 18),
                    "average_salary": base_salary
                }
            }
            colleges.append(college)
            college_id += 1
    
    # Continue until 1300
    while len(colleges) < 1300:
        province_name, capital, _ = provinces[(len(colleges) - 40) % len(provinces)]
        cities = cities_by_province.get(province_name, [capital])
        city = cities[len(colleges) % len(cities)]
        college_type = types[len(colleges) % len(types)]
        
        base_rank = len(colleges) + 1
        base_salary = 5500 + (len(colleges) % 35) * 100
        base_enrollment = 1800 + (len(colleges) % 25) * 80
        
        suffixes = ["University", "College", "Institute", "Vocational College", "Applied University", "Tech University"]
        name = f"{city} {suffixes[len(colleges) % len(suffixes)]}"
        if len(colleges) % 3 == 0:
            name = f"{province_name} {suffixes[len(colleges) % len(suffixes)]}"
        
        college = {
            "id": f"{len(colleges) + 1:05d}",
            "name": name,
            "province": province_name,
            "city": city,
            "type": college_type,
            "level": "Undergraduate",
            "established_year": 1978 + (len(colleges) % 45),
            "departments": 8 + (len(colleges) % 12),
            "faculties": 3500 + (len(colleges) % 30) * 100,
            "website": f"https://www.edu{len(colleges) + 1}.cn",
            "description": f"University in {province_name}",
            "rankings": {
                "comprehensive": base_rank,
                "qs_world": base_rank * 10
            },
            "majors": get_majors_for_college(college_type, base_rank, city, province_name),
            "statistics": {
                "enrollment_count": base_enrollment,
                "employment_rate": 78 + (len(colleges) % 17),
                "average_salary": base_salary
            }
        }
        colleges.append(college)
    
    return colleges

def generate_majors():
    majors = []
    major_id = 1
    for name, info in MAJORS_DB.items():
        majors.append({
            "id": f"{(major_id):05d}",
            "name": name,
            "category": info["category"],
            "code": info["code"],
            "degree": f"Bachelor of {info['category']}",
            "duration": info.get("duration", 4),
            "suitable_types": info["types"],
            "employment": {
                "rate": round(85 + random.random() * 12, 1),
                "avg_salary": round(7000 + random.random() * 10000)
            }
        })
        major_id += 1
    return majors

def generate_scores(colleges, majors):
    scores = []
    score_id = 1
    
    for college in colleges[:600]:
        college_majors = college.get("majors", [])[:8]
        if not college_majors:
            continue
            
        for major in college_majors:
            for year in [2024, 2023, 2022]:
                for category in ["Science", "Liberal Arts"]:
                    rank = college["rankings"]["comprehensive"]
                    if rank < 50:
                        base_score = 680 - rank * 3
                    elif rank < 100:
                        base_score = 580 - (rank - 50) * 2
                    elif rank < 200:
                        base_score = 520 - (rank - 100)
                    elif rank < 500:
                        base_score = 480 - (rank - 200) * 0.5
                    else:
                        base_score = 420 - (rank - 500) * 0.2
                    
                    if category == "Liberal Arts":
                        base_score -= 35
                    
                    score = int(base_score + (score_id % 25))
                    rank_min = 100 + score_id * 50
                    rank_max = rank_min + 800
                    
                    scores.append({
                        "id": f"{30000 + score_id}",
                        "college_id": college["id"],
                        "college_name": college["name"],
                        "major_name": major["name"],
                        "year": year,
                        "batch": "Batch 1" if rank < 200 else "Batch 2",
                        "category": category,
                        "score": score,
                        "rank_min": rank_min,
                        "rank_max": rank_max
                    })
                    score_id += 1
                    if score_id > 20000:
                        break
                if score_id > 20000:
                    break
            if score_id > 20000:
                break
        if score_id > 20000:
            break
    
    return scores

def generate_rankings(colleges):
    rankings = []
    rank_id = 1
    
    for college in colleges:
        rankings.append({
            "id": f"{40000 + rank_id}",
            "year": 2024,
            "ranking_type": "comprehensive",
            "college_id": college["id"],
            "college_name": college["name"],
            "rank": college["rankings"]["comprehensive"]
        })
        rank_id += 1
        
        if college["rankings"]["comprehensive"] <= 200:
            for subject in ["engineering", "science", "medicine", "social_science"]:
                rankings.append({
                    "id": f"{50000 + rank_id}",
                    "year": 2024,
                    "ranking_type": f"subject_{subject}",
                    "college_id": college["id"],
                    "college_name": college["name"],
                    "rank": college["rankings"]["comprehensive"] + (rank_id % 30)
                })
                rank_id += 1
    
    return rankings

def main():
    print('='*60)
    print('Gaokao Advisor - Complete Data Collection')
    print('Including College-Major Associations')
    print('='*60)
    
    print("\n[1/5] Generating college data with majors...")
    colleges = generate_colleges()
    print(f"  Colleges: {len(colleges)}")
    
    print("\n[2/5] Generating major data...")
    majors = generate_majors()
    print(f"  Majors: {len(majors)}")
    
    print("\n[3/5] Generating college-major associations...")
    college_majors = []
    cm_id = 1
    for college in colleges:
        for major in college.get("majors", []):
            college_majors.append({
                "id": f"{cm_id:06d}",
                "college_id": college["id"],
                "college_name": college["name"],
                "major_name": major["name"],
                "major_code": major["code"],
                "category": major["category"],
                "degree": major["degree"],
                "duration": major["duration"]
            })
            cm_id += 1
    print(f"  College-Major links: {len(college_majors)}")
    
    print("\n[4/5] Generating admission score data...")
    scores = generate_scores(colleges, majors)
    print(f"  Scores: {len(scores)}")
    
    print("\n[5/5] Generating ranking data...")
    rankings = generate_rankings(colleges)
    print(f"  Rankings: {len(rankings)}")
    
    print("\n[6/6] Saving data to files...")
    save_json(colleges, 'colleges.json')
    save_json(majors, 'majors.json')
    save_json(college_majors, 'college_majors.json')
    save_json(scores, 'scores.json')
    save_json(rankings, 'rankings.json')
    
    print("\n" + "="*60)
    print("Data collection completed!")
    print("="*60)
    print(f"\nSummary:")
    print(f"  - Colleges: {len(colleges)}")
    print(f"  - Majors: {len(majors)}")
    print(f"  - College-Major Links: {len(college_majors)}")
    print(f"  - Admission Scores: {len(scores)}")
    print(f"  - Rankings: {len(rankings)}")
    
    # Sample
    print("\nSample - Tsinghua University majors:")
    for m in colleges[0].get("majors", [])[:5]:
        print(f"  - {m['name']} ({m['code']})")

if __name__ == '__main__':
    random.seed(42)
    main()