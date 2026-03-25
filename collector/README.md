# Data Collector Service

## Overview

This Python service handles web scraping and data collection for the Gaokao Advisory System.

## Data Sources

- 阳光高考平台 (https://gaokao.chsi.com.cn)
- 中国教育在线 (https://www.eol.cn)
- 河南省教育考试院 (https://www.haeea.cn)

## Installation

```bash
pip install -r requirements.txt
```

## Usage

```bash
python run.py --help
```

## Commands

- `python run.py colleges` - Collect college data
- `python run.py majors` - Collect major data  
- `python run.py scores` - Collect admission scores
- `python run.py all` - Collect all data

## Configuration

Edit `config.py` to set target data sources and output paths.
