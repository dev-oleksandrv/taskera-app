"use client";

import { PropsWithChildren, useEffect, useState } from "react";
import { Sidebar } from "@/components/dashboard/sidebar";
import { cn } from "@/lib/utils";
import { Header } from "@/components/dashboard/header";
import { ListDto, SpaceWithRoleDto } from "@/types/server-api-types";
import { useSpaceStore } from "@/store/spaceStore";

interface ContainerProps {
  spaces: SpaceWithRoleDto[];
  lists: ListDto[];

  currentSpaceID: string;
  currentListID: string;
}

export const Container = ({ children, spaces, lists, currentSpaceID, currentListID }: PropsWithChildren<ContainerProps>) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  const toggleSidebar = () => setIsSidebarOpen((prevState) => !prevState);

  useEffect(() => {
    useSpaceStore.setState({
      spaces,
      currentSpaceID,
      lists,
      currentListID
    });
  }, []);

  return (
    <>
      <Sidebar sidebarOpen={isSidebarOpen} onSidebarToggle={toggleSidebar} />

      <main
        className={cn(
          "flex-1 overflow-hidden",
          isSidebarOpen ? "ml-56" : "ml-12"
        )}
      >
        <Header />
        {children}
      </main>
    </>
  )
}