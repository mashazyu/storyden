import { LinkList } from "src/api/openapi/schemas";
import { LinkRef } from "src/components/directory/links/LinkRef";
import { Empty } from "src/components/feed/common/PostRef/Empty";

import { styled } from "@/styled-system/jsx";

export function LinkRefList(props: { links: LinkList }) {
  if (props.links.length === 0) {
    return <Empty>no links were found</Empty>;
  }

  return (
    <styled.ol display="flex" flexDir="column" gap="4">
      {props.links.map((v) => (
        <styled.li key={v.url}>
          <LinkRef {...v} />
        </styled.li>
      ))}
    </styled.ol>
  );
}