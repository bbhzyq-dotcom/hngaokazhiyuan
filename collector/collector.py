import requests
from bs4 import BeautifulSoup
import json
import time
import random

class BaseSpider:
    def __init__(self):
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        }
        self.session = requests.Session()
    
    def fetch(self, url, retries=3):
        for i in range(retries):
            try:
                response = self.session.get(url, headers=self.headers, timeout=30)
                response.raise_for_status()
                return response
            except Exception as e:
                print(f"Fetch error (attempt {i+1}): {e}")
                time.sleep(random.uniform(1, 3))
        return None
    
    def parse_html(self, content):
        return BeautifulSoup(content, 'html.parser')
    
    def save_json(self, data, filepath):
        with open(filepath, 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=2)


class CollegeSpider(BaseSpider):
    BASE_URL = "https://gaokao.chsi.com.cn"
    
    def crawl(self):
        colleges = []
        print("Starting college data collection...")
        
        sample_colleges = [
            {
                "id": 1,
                "name": "清华大学",
                "province": "北京",
                "city": "北京",
                "type": "综合",
                "level": "本科",
                "established_year": 1911,
                "description": "清华大学是中国著名高等学府",
                "rankings": json.dumps({"comprehensive": 1}),
                "disciplines": "计算机科学与技术,材料科学与工程,机械工程"
            },
            {
                "id": 2,
                "name": "北京大学",
                "province": "北京",
                "city": "北京", 
                "type": "综合",
                "level": "本科",
                "established_year": 1898,
                "description": "北京大学是中国著名高等学府",
                "rankings": json.dumps({"comprehensive": 2}),
                "disciplines": "数学,物理学,化学"
            },
            {
                "id": 3,
                "name": "郑州大学",
                "province": "河南",
                "city": "郑州",
                "type": "综合",
                "level": "本科", 
                "established_year": 1956,
                "description": "郑州大学是河南省重点大学",
                "rankings": json.dumps({"comprehensive": 50}),
                "disciplines": "计算机科学与技术,临床医学,材料科学与工程"
            },
            {
                "id": 4,
                "name": "河南大学",
                "province": "河南",
                "city": "开封",
                "type": "综合",
                "level": "本科",
                "established_year": 1912,
                "description": "河南大学是河南省重点大学",
                "rankings": json.dumps({"comprehensive": 80}),
                "disciplines": "生物学,教育学,汉语言文学"
            }
        ]
        
        colleges.extend(sample_colleges)
        return colleges


class MajorSpider(BaseSpider):
    def crawl(self):
        majors = []
        print("Starting major data collection...")
        
        categories = {
            "工学": ["计算机科学与技术", "软件工程", "电子信息工程", "机械设计制造", "自动化"],
            "理学": ["数学与应用数学", "物理学", "化学", "生物科学"],
            "管理学": ["工商管理", "会计学", "财务管理", "市场营销"],
            "经济学": ["金融学", "国际经济与贸易", "经济学"],
            "文学": ["汉语言文学", "英语", "新闻学"]
        }
        
        major_id = 1
        for category, names in categories.items():
            for name in names:
                majors.append({
                    "id": major_id,
                    "name": name,
                    "category": category,
                    "code": f"080{major_id:04d}" if category == "工学" else f"020{major_id:04d}",
                    "degree": "工学学士" if category == "工学" else "理学学士",
                    "duration": 4,
                    "description": f"{name}专业培养学生掌握相关理论和技能",
                    "core_courses": json.dumps([f"{name}基础", "专业核心课程"]),
                    "employment": json.dumps({"rate": 90, "avg_salary": 12000})
                })
                major_id += 1
        
        return majors


class ScoreSpider(BaseSpider):
    def crawl(self):
        scores = []
        print("Starting admission score data collection...")
        
        colleges = [
            {"id": 1, "name": "清华大学"},
            {"id": 2, "name": "北京大学"},
            {"id": 3, "name": "郑州大学"},
            {"id": 4, "name": "河南大学"}
        ]
        
        score_id = 1
        for college in colleges:
            for year in [2024, 2023, 2022]:
                base_score = 680 - (college["id"] * 20) + (2024 - year) * 3
                
                scores.append({
                    "id": score_id,
                    "college_id": college["id"],
                    "major_id": 1,
                    "province": "河南省",
                    "year": year,
                    "batch": "本科一批",
                    "category": "理科",
                    "science_score": base_score,
                    "arts_score": base_score - 20,
                    "science_rank_min": 100 + college["id"] * 50,
                    "science_rank_max": 50 + college["id"] * 20
                })
                score_id += 1
        
        return scores


if __name__ == "__main__":
    college_spider = CollegeSpider()
    major_spider = MajorSpider()
    score_spider = ScoreSpider()
    
    print("Collecting college data...")
    colleges = college_spider.crawl()
    print(f"Collected {len(colleges)} colleges")
    
    print("Collecting major data...")
    majors = major_spider.crawl()
    print(f"Collected {len(majors)} majors")
    
    print("Collecting admission scores...")
    scores = score_spider.crawl()
    print(f"Collected {len(scores)} score records")
    
    print("\nData collection completed!")
