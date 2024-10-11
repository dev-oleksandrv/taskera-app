"use server";

import { getAllTasksBySpaceListIDRequest } from "@/lib/api/server-api";

export const getTasksAction = async (spaceID: string, listID: string) => {
  const response =  await getAllTasksBySpaceListIDRequest(spaceID, listID);

  return response.data.tasks;
}