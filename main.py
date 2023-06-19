import csv
import requests
from bs4 import BeautifulSoup

def main():
    url = ""

    response = requests.get(url)
    response.raise_for_status()

    soup = BeautifulSoup(response.content, "html.parser")

    with open("dadosPy.csv", "w", newline="", encoding="utf-8") as csvfile:
        writer = csv.writer(csvfile)

        for row in soup.select("table tr"):
            row_data = []

            for col in row.select("th, td"):
                cell_data = col.get_text().strip()

                if not contains_ignored_keyword(cell_data):
                    row_data.append(cell_data)

            if row_data:
                writer.writerow(row_data)

    print("Dados gravados no arquivo dados.csv.")

def contains_ignored_keyword(text):
    ignored_keywords = ["DIAG", "CONF", "TOOL", "CMD"]
    for keyword in ignored_keywords:
        if keyword in text:
            return True
    return False

if __name__ == "__main__":
    main()
