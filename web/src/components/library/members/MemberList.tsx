import { PublicProfileList } from "src/api/openapi-schema";

import { Empty } from "src/components/site/Empty";
import { MemberCard } from "./MemberCard";
import * as Table from "src/components/ui/table";

type Props = {
  profiles: PublicProfileList;
  onChange: () => void;
};

export function MemberList(props: Props) {
  if (props.profiles.length === 0) {
    return <Empty>no members were found</Empty>;
  }

  return (
    <Table.Root>
    <Table.Caption>Recent members</Table.Caption>
    <Table.Head>
      <Table.Row textTransform="uppercase">
        <Table.Header>Name</Table.Header>
        <Table.Header>Member since</Table.Header>
        <Table.Header>Roles</Table.Header>
      </Table.Row>
    </Table.Head>
    <Table.Body>
      {props.profiles.map((profile) => (
        < MemberCard profile={profile} />
      ))}
    </Table.Body>
  </Table.Root>
  );
}
