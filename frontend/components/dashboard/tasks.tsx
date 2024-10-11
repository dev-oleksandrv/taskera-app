"use client";

import { useEffect, useMemo } from "react";
import { Task } from "@/components/dashboard/task";
import { getTasksAction } from "@/app/dashboard/[spaceID]/[listID]/actions";
import { useSpaceStore } from "@/store/spaceStore";
import { useTaskStore } from "@/store/taskStore";
import { TaskDto } from "@/types/server-api-types";

export const Tasks = () => {
  const { currentSpaceID, currentListID } = useSpaceStore();
  const { tasks, setTasks } = useTaskStore();

  const [ongoing, completed] = useMemo(() => {
    const [ongoing, completed] = [[], []] as [TaskDto[], TaskDto[]];
    tasks.forEach(t => {
      if (t.completed) {
        completed.push(t)
      } else {
        ongoing.push(t)
      }
    })
    return [ongoing, completed]
  }, [tasks])

  useEffect(() => {
    if (!currentSpaceID || !currentListID) return;

    getTasksAction(currentSpaceID, currentListID).then(tasks => {
      setTasks(tasks)
    })
  }, [currentSpaceID, currentListID]);

  return (
    <div className="space-y-2 mb-4">
      <h2 className="font-semibold text-gray-700 mb-2">Ongoing</h2>

      {ongoing.map(task =>
        <Task key={task.id} task={task} />
      )}

      <h2 className="font-semibold text-gray-700 mb-2">Completed</h2>

      {completed.map(task =>
        <Task key={task.id} task={task} />
      )}
    </div>
  )
}
