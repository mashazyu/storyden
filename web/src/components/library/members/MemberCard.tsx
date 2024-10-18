import { formatDistanceToNow } from "date-fns";

import { PublicProfile } from "src/api/openapi-schema";

import * as Table from "src/components/ui/table";
import { MemberBadge } from "@/components/member/MemberBadge/MemberBadge";
import { MemberOptionsTrigger } from "@/components/member/MemberOptions/MemberOptionsTrigger";
import { Avatar } from "@/components/site/Avatar/Avatar";
import { RoleBadge } from "@/components/member/RoleBadge/RoleBadge";

import { Box, LinkOverlay, VStack, styled } from "@/styled-system/jsx";

type Props = { profile: PublicProfile };

export function MemberCard(props: Props) {
  const { profile } = props; 
  
  return (
    <Table.Row key={profile.id}>
          <Table.Cell display="flex" gap="2" alignItems="center">
            <Avatar handle={profile.handle} />
            <VStack alignItems="start" gap="0">
              <styled.h1 color="fg.default">{profile.name}</styled.h1>
              <styled.h2 color="fg.subtle">@{profile.handle}</styled.h2>
            </VStack>
          </Table.Cell>
          <Table.Cell>{formatDistanceToNow(profile.createdAt)}</Table.Cell>
          <Table.Cell>{profile.roles.map((role) => <RoleBadge role={role} />)}</Table.Cell>
          <Table.Cell textAlign="right"><MemberOptionsTrigger {...profile} /></Table.Cell>
        </Table.Row>
  );
}
