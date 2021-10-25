import yaml
import os
from spider import GitRepoSpider
from download import DownloadManager


class GitRepoManager:

    def __init__(self):
        current_path = os.path.abspath(os.path.dirname(__file__))
        self.config_path = os.path.join(current_path, "config.yaml")
        config_file = open(self.config_path, 'r', encoding='utf-8').read()
        self.config = yaml.load(config_file, Loader=yaml.BaseLoader)
        self.proxy = self.config.get('proxy')
        self.path = self.config.get('path')

    def check_update(self):
        print("")
        up = True
        for repo in self.config['repos']:
            spider = GitRepoSpider(repo['name'], repo['author'], repo['tag'], repo['key'], self.proxy)
            spider.get_latest_tag()
            if spider.flag:
                up = False
        if up:
            print("Everything is up to date.")
        print("")

    def update_all(self):
        for repo in self.config['repos']:
            spider = GitRepoSpider(repo['name'], repo['author'], repo['tag'], repo['key'], self.proxy)
            tag = spider.get_latest_tag()
            if spider.flag:
                spider.get_assets()
                downloading = DownloadManager(repo['name'], self.path, tag, self.proxy)
                print("Start downloading from "+spider.download_url)
                downloading.download(spider.download_url, spider.download_name)
        print("Now everything is up to date.\n")

    def list_all(self):
        print("\nInstalled apps: \n")
        for repo in self.config['repos']:
            print(repo['name'] + " " + repo['tag'] + "["+repo['author']+"]")
        print("")

    def modify_config(self):
        try:
            os.system('code '+self.config_path)
        except:
            os.system('notepad ' + self.config_path)
        print("")