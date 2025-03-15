import sys

import os
import pandas as pd


def parse_excel_to_txt(excel_file, vocabulary):
    # 读取 Excel 文件
    try:
        # 读取所有 Sheet，找到"高中词汇"
        xls = pd.ExcelFile(excel_file)

        if "高中词汇" not in xls.sheet_names:
            raise ValueError(f"未找到名为 {vocabulary} 的 Sheet")

        # 读取 Sheet
        df = pd.read_excel(excel_file, sheet_name=vocabulary)
    except Exception as e:
        print(f"读取 Excel 文件失败: {e}")
        return

    # 确保 目录, 英文单词, 中文意思 列存在
    if df.shape[1] < 3:  # df.shape[1] 可以获取所有的列
        print("Excel 文件至少需要 3 列（目录, 英文单词, 中文意思）")
        return

    # 按 目录 列分组
    grouped = df.groupby('目录')

    # 创建目录（如果不存在）
    output_dir = vocabulary
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)
        print(f"已创建目录: {output_dir}")

    # 遍历每个分组，生成 txt 文件
    for list_name, group in grouped:
        # 确保 list_name 是字符串，移除可能的空格
        list_name = str(list_name).strip()
        if not list_name:
            continue  # 跳过空的 list_name

        # 创建 txt 文件，文件名如 list1.txt
        txt_filename = os.path.join(output_dir, f"{list_name}.txt")
        try:
            with open(txt_filename, 'w', encoding='utf-8') as f:
                # 遍历分组中的每一行
                for index, row in group.iterrows():
                    word = str(row['英文单词']).strip()
                    meaning = str(row['中文意思']).strip()
                    if word and meaning:  # 确保单词和意思不为空
                        f.write(f"{word} | {meaning}\n")
            print(f"已生成 {txt_filename}")
        except Exception as e:
            print(f"写入 {txt_filename} 失败: {e}")


if __name__ == "__main__":
    # 替换为你的 Excel 文件路径
    excel_file_path = "单词收录.xlsx"

    # 获取命令行参数
    if len(sys.argv) != 2:
        print("用法: python parse-excel-2-txt.py <output_dir_name>")
        print("示例: python parse-excel-2-txt.py 高中词汇")
        sys.exit(1)

    vocabulary = sys.argv[1]  # 参数是哪个单词
    parse_excel_to_txt(excel_file_path, vocabulary)
