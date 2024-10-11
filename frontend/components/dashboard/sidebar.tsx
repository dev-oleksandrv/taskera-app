import Link from "next/link";
import { Button } from "@/components/ui/button";
import { ChevronLeftIcon, ChevronRightIcon } from "lucide-react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { ScrollArea } from "@/components/ui/scroll-area";
import { useSpaceStore } from "@/store/spaceStore";
import { useRouter } from "next/navigation";

interface SidebarProps {
  sidebarOpen: boolean;
  onSidebarToggle: () => void;
}

export const Sidebar = ({ sidebarOpen, onSidebarToggle }: SidebarProps) => {
  const { spaces, lists, currentSpaceID, currentListID, setSpace } = useSpaceStore()

  const router = useRouter();

  const onSpaceChange = (value: string) => {
    router.replace(`/dashboard/${value}/~`);
    setSpace(value);
  }

  const onListChange = (value: string) => {
    router.replace(`/dashboard/${currentSpaceID}/${value}`)
  }

  return (
    <aside
      className={`fixed left-0 top-0 bottom-0 bg-white border-r transition-all duration-300 ${
        sidebarOpen ? "w-56" : "w-12"
      } flex flex-col z-20`}
    >
      <div className="p-2 flex items-center justify-between">
        {sidebarOpen && (
          <Link href="/" className="flex items-center">
            <span className="text-xl font-bold text-teal-700">Taskera</span>
          </Link>
        )}
        <Button
          variant="ghost"
          size="icon"
          onClick={onSidebarToggle}
          className="ml-auto"
        >
          {sidebarOpen ? (
            <ChevronLeftIcon className="h-4 w-4" />
          ) : (
            <ChevronRightIcon className="h-4 w-4" />
          )}
        </Button>
      </div>
      {sidebarOpen ? (
        <>
          <div className="p-2">
            <Select
              disabled={false}
              value={currentSpaceID}
              onValueChange={onSpaceChange}
            >
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select space" />
              </SelectTrigger>
              <SelectContent>
                {spaces.map((space) => (
                  <SelectItem key={space.space.id} value={`${space.space.id}`}>
                    {space.space.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
          <ScrollArea className="flex-1">
            <nav className="space-y-1 p-2">
              {lists.map((list) => (
                <Button
                  key={list.id}
                  variant={currentListID === list.id ? "secondary" : "ghost"}
                  className="w-full justify-start"
                  onClick={() => onListChange(list.id)}
                >
                  <span className="mr-2">{list.emoji}</span>
                  {list.name}
                </Button>
              ))}
            </nav>
          </ScrollArea>
        </>
      ) : (
        <ScrollArea className="flex-1">
          <nav className="space-y-1 p-2">
            {lists.map((list) => (
              <Button
                key={list.id}
                variant={currentListID === list.id ? "secondary" : "ghost"}
                size="icon"
                className="w-full"
              >
                <span>{list.emoji}</span>
              </Button>
            ))}
          </nav>
        </ScrollArea>
      )}
    </aside>
  )
}