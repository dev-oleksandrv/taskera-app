import { getAllListsBySpaceIDRequest, getAllSpacesRequest } from "@/lib/api/server-api";
import { Container } from "@/components/dashboard/container";
import { redirect } from "next/navigation";
import { SpaceWithRoleDto } from "@/types/server-api-types";
import { Tasks } from "@/components/dashboard/tasks";

export default async function DashboardListPage({ params: { spaceID, listID } }: {
  params: {
    spaceID: string;
    listID: string;
  }
}) {
  const spaceResponse = await getAllSpacesRequest();
  if (!spaceResponse.data.spaces.length) {
    return redirect("/dashboard/create-first-space");
  }

  let currentSpace = spaceResponse.data.spaces.find(sp => sp.space.id === spaceID);
  if (!currentSpace) {
    currentSpace = spaceResponse.data.spaces[0] as SpaceWithRoleDto;
  }

  const listResponse = await getAllListsBySpaceIDRequest(currentSpace.space.id);
  if (!listResponse.data.lists.length) {
    return redirect(`/dashboard/${currentSpace.space.id}`)
  }

  let currentList = listResponse.data.lists.find(lt => lt.id === listID);
  if (!currentList) {
    return redirect(`/dashboard/${currentSpace.space.id}/${listResponse.data.lists[0].id}`)
  }

  return (
    <div className="flex h-screen bg-gray-50">
      <Container lists={listResponse.data.lists} spaces={spaceResponse.data.spaces} currentListID={currentList.id} currentSpaceID={currentSpace.space.id}>
        <div className="p-4 h-[calc(100vh-3rem)] overflow-y-auto">
          <Tasks />
        </div>
      </Container>
    </div>
  )
}