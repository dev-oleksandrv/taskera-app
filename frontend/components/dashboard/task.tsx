import { GripVerticalIcon } from "lucide-react";
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";
import { TaskDto } from "@/types/server-api-types";
import { useTaskStore } from "@/store/taskStore";

interface TaskProps {
  task: TaskDto
}

export const Task = ({ task }: TaskProps) => {
  const { toggleTask } = useTaskStore();

  const onToggle = (v: boolean) => toggleTask(task.id, v)

  return (
    <div
      className="bg-white rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-shadow duration-200 flex items-center group">
      <div className="p-2 cursor-move">
        <GripVerticalIcon className="h-4 w-4 text-gray-400" />
      </div>
      <Checkbox
        className="ml-2"
        checked={task.completed}
        onCheckedChange={onToggle}
      />
      <Button variant="ghost" className="flex-grow text-left justify-start p-2 h-auto">
        <span>{task.content}</span>
      </Button>
    </div>
  )
}