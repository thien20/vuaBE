import os

data_folder = r'D:\MY_FOLDER\Project\vuaBE\data'
all_files = [f for f in os.listdir(data_folder) if os.path.isfile(os.path.join(data_folder, f))]
print(all_files)