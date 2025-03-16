import sys

from datetime import datetime

abbreviation_dict = {
    'cz': '初中词汇',
    'gz': '高中词汇',
    'sj': '四级词汇',
    'lj': '六级词汇',
}


def main(filename, words_book, chapter):
    # 读取文件并处理空行
    words = []
    origin_words = []
    origin_words_map = {}
    mismatched = []

    if words_book not in abbreviation_dict:
        print(f"{words_book}: 没有找到匹配的词汇书")
        return

    book_name = abbreviation_dict[words_book]
    with open(filename, 'r', encoding='utf-8') as file:
        # 使用列表推导式，strip()去掉首尾空白，过滤空行
        words = [line.strip() for line in file if line.strip()]

    # print(words)

    with open(f"{book_name}/{chapter}.txt", 'r', encoding='utf-8') as file:
        # 逐行读取
        for line in file:
            # 去除首尾空白
            line = line.strip()
            # 如果行不为空
            if line:
                # 用 '|' 分割，取第一部分（即 '|' 前的单词）
                word = line.split('|')[0].strip()
                meaning = line.split('|')[1].strip()
                origin_words.append(word)
                origin_words_map[word] = meaning
    # print(origin_words)

    # 确保两个列表长度相同
    if len(words) != len(origin_words):
        print(f"警告：列表长度不一致，听写的单词长度为: {len(words)}, 原列表单词个数是: {len(origin_words)}")
        return

    # 逐个比较
    words_length = len(origin_words)
    for i in range(words_length):
        # 不区分大小写比较
        if words[i].lower() != origin_words[i].lower():
            mismatched.append((words[i], origin_words[i]))

    # 输出结果
    if mismatched:
        print(f"""
单词打卡。
{book_name}/{chapter} 共有单词: {len(origin_words)}, 共错误了 {len(mismatched)} 个单词。听写时间：{datetime.now().strftime("%Y-%m-%d %H:%M:%S")}。
错误的单词如下，错误单词 -> 正确单词
""")
        for wrong, correct in mismatched:
            print(f"{wrong} -> {correct}")
    else:
        print("恭喜你，所有单词听写正确！")

    # 将错误单词写入文件
    if mismatched:
        try:
            wrong_words_filename = f"wrong_words/{datetime.now().strftime("%Y-%m-%d")}-{book_name}-{chapter}.txt"
            with open(wrong_words_filename, 'w', encoding='utf-8') as f:
                for wrong, correct in mismatched:
                    # 写入格式：correct | definition
                    f.write(f"{correct} | {origin_words_map[correct]}\n")
            print(f"\n错误单词已写入到 {wrong_words_filename} 文件中，方便复习。")
        except Exception as e:
            print(f"写入文件时发生错误：{e}")

    return


if __name__ == "__main__":
    # 获取命令行参数
    if len(sys.argv) != 4:
        print("""
简称如下：
        
    cz -> 初中词汇
    gz -> 高中词汇
    sj -> 四级词汇
    lj -> 六级词汇
    
使用方法：

    python  check.py dictation/xxx.txt sj 具体哪个章节
    比如：
    python check.py  dictation\20250315.txt gz list1
""")
        sys.exit(1)

    filename = sys.argv[1]  # 听写的词汇
    words_book = sys.argv[2]  # 词汇书
    chapter = sys.argv[3]
    main(filename, words_book, chapter)
