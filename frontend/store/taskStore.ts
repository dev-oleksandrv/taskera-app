import { TaskDto } from "@/types/server-api-types";
import { create } from "zustand"

export interface TaskStore {
  tasks: TaskDto[];
  setTasks: (tasks: TaskDto[]) => void;

  toggleTask: (taskID: string, completeStatus: boolean) => void;
}

export const useTaskStore = create<TaskStore>((set, getState) => ({
  tasks: [],
  setTasks: (tasks: TaskDto[]) => set({ tasks }),

  toggleTask: (taskID: string, completeStatus: boolean) => set({
    tasks: getState().tasks.map(t => {
      if (t.id === taskID) {
        return {
          ...t,
          completed: completeStatus
        }
      }
      return t;
    })
  })
}))