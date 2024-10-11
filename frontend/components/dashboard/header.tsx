import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { BellIcon, LogOutIcon, SettingsIcon } from "lucide-react";
import { ScrollArea } from "@/components/ui/scroll-area";
import { ChevronDownIcon } from "@radix-ui/react-icons";
import { useSpaceStore } from "@/store/spaceStore";

const onlineUsers = [
  { id: 1, name: "John Doe", avatar: "/john-doe.jpg" },
  { id: 2, name: "Jane Smith", avatar: "/jane-smith.jpg" },
  { id: 3, name: "Alice Johnson", avatar: "/alice-johnson.jpg" },
];

export const Header = () => {
  const selectedList = useSpaceStore((state) => {
    if (!state.currentListID || !state.lists.length) return null;

    return state.lists.find((item) => item.id === state.currentListID) || null;
  });

  return (
    <header className="bg-white border-b h-12 flex items-center px-4">
      {selectedList && (
        <h1 className="font-semibold text-teal-700 mr-4">
          {selectedList.emoji} {selectedList.name}
        </h1>
      )}
      <div className="flex items-center space-x-2 ml-auto">
        <div className="flex -space-x-2">
          {onlineUsers.map((user) => (
            <Avatar key={user.id} className="border-2 border-white h-8 w-8">
              <AvatarImage src={user.avatar} alt={user.name} />
              <AvatarFallback>{user.name.charAt(0)}</AvatarFallback>
            </Avatar>
          ))}
        </div>
        <Popover>
          <PopoverTrigger asChild>
            <Button variant="ghost" size="icon" className="h-8 w-8">
              <BellIcon className="h-4 w-4" />
            </Button>
          </PopoverTrigger>
          <PopoverContent className="w-80 p-0">
            <div className="p-2 border-b">
              <h3 className="font-semibold">Notifications</h3>
            </div>
            <ScrollArea className="h-[300px]">
              <div className="space-y-2 p-2">
                <div className="flex items-start space-x-2">
                  <Avatar className="h-8 w-8">
                    <AvatarImage src="/jane-smith.jpg" alt="Jane Smith" />
                    <AvatarFallback>JS</AvatarFallback>
                  </Avatar>
                  <div>
                    <p className="text-sm">
                      Jane Smith made changes in Work/To-Do
                    </p>
                    <p className="text-xs text-gray-500">2 hours ago</p>
                  </div>
                </div>
                <div className="flex items-start space-x-2">
                  <Avatar className="h-8 w-8">
                    <AvatarImage src="/john-doe.jpg" alt="John Doe" />
                    <AvatarFallback>JD</AvatarFallback>
                  </Avatar>
                  <div>
                    <p className="text-sm">
                      John Doe completed a task in Personal/In Progress
                    </p>
                    <p className="text-xs text-gray-500">5 hours ago</p>
                  </div>
                </div>
              </div>
            </ScrollArea>
          </PopoverContent>
        </Popover>
        <Popover>
          <PopoverTrigger asChild>
            <Button variant="ghost" size="sm" className="h-8">
              <Avatar className="h-6 w-6">
                <AvatarImage src="/placeholder-avatar.jpg" alt="User" />
                <AvatarFallback>JD</AvatarFallback>
              </Avatar>
              <ChevronDownIcon className="h-4 w-4 ml-2" />
            </Button>
          </PopoverTrigger>
          <PopoverContent className="w-56 p-0">
            <div className="space-y-1">
              <Button
                variant="ghost"
                size="sm"
                className="w-full justify-start"
              >
                <SettingsIcon className="mr-2 h-4 w-4" />
                Settings
              </Button>
              <Button
                variant="ghost"
                size="sm"
                className="w-full justify-start text-red-500"
              >
                <LogOutIcon className="mr-2 h-4 w-4" />
                Logout
              </Button>
            </div>
          </PopoverContent>
        </Popover>
      </div>
    </header>
  );
};