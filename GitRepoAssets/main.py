from manager import GitRepoManager

if __name__ == '__main__':
    manager = GitRepoManager()
    while True:
        print(
            "###################################\n"
            "Welcome to GitRepoAssets Manager\n"
            "Enter help for more actions\n\n"
            "Your command here:"
        )

        command = input()
        if command == "help":
            print(
                "\nUsage: <command> [<args>]\n\n"
                "Some useful commands are:\n\n"
                "list/ls                :           List installed apps\n"
                "status/st              :           Show status and check for new app versions\n"
                "update/up              :           Update all apps\n"
                "config/cf              :           Open config file to add an app or modify other settings\n"
                "exit/et                :           Exit the shell\n"
            )
        elif command == "list" or command == "ls":
            manager.list_all()
        elif command == "status" or command == "st":
            manager.check_update()
        elif command == "update" or command =="up":
            manager.update_all()
        elif command == "config" or command =="cf":
            manager.modify_config()
        elif command == "exit" or command =="et":
            break
        else:
            print("\nCommand not defined\n")
