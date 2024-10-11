import { ListDto, SpaceWithRoleDto } from "@/types/server-api-types";
import { create } from "zustand"

export interface SpaceStore {
  spaces: SpaceWithRoleDto[];
  currentSpaceID: string;
  setSpace: (spaceID: string) => void;

  lists: ListDto[];
  currentListID: string;
}

export const useSpaceStore = create<SpaceStore>(set => ({
  spaces: [],
  currentSpaceID: "",
  setSpace: (spaceID: string) => set({ currentSpaceID: spaceID }),

  lists: [],
  currentListID: "",
}))