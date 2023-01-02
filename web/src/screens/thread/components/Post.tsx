import {
  Code,
  Flex,
  Heading,
  Link,
  ListItem,
  OrderedList,
  Text,
  UnorderedList,
} from "@chakra-ui/react";
import ReactMarkdown from "react-markdown";
import { SpecialComponents } from "react-markdown/lib/ast-to-react";
import { NormalComponents } from "react-markdown/lib/complex-types";
import { Post } from "src/api/openapi/schemas";
import { Byline } from "./Byline";

const components: Partial<
  Omit<NormalComponents, keyof SpecialComponents> & SpecialComponents
> = {
  h1: (props) => (
    <Heading as="h1" size="2xl">
      {props.children}
    </Heading>
  ),
  h2: (props) => (
    <Heading as="h2" size="xl">
      {props.children}
    </Heading>
  ),
  h3: (props) => (
    <Heading as="h3" size="lg">
      {props.children}
    </Heading>
  ),
  h4: (props) => (
    <Heading as="h4" size="md">
      {props.children}
    </Heading>
  ),
  h5: (props) => (
    <Heading as="h5" size="sm">
      {props.children}
    </Heading>
  ),

  // Typography
  p: (props) => (
    <Text overflowWrap="anywhere" wordBreak="break-all" overflowX="clip">
      {props.children}
    </Text>
  ),
  a: ({ href, children }) => <Link href={href ?? "#"}>{children}</Link>,

  // Lists
  ul: (props) => <UnorderedList ml="2em">{props.children}</UnorderedList>,
  ol: (props) => <OrderedList ml="2em">{props.children}</OrderedList>,
  li: (props) => <ListItem>{props.children}</ListItem>,

  // Code
  pre: (props) => (
    <Code
      display="block"
      whiteSpace="pre"
      overflowX="scroll"
      padding={2}
      borderRadius="md"
    >
      {props.children}
    </Code>
  ),
  code: (props) => <Code>{props.children}</Code>,

  // Headings
  td: (props) => <td>{props.children}</td>,
  th: (props) => <th>{props.children}</th>,
  tr: (props) => <tr>{props.children}</tr>,
};
export function PostView(props: Post) {
  return (
    <Flex flexDir="column">
      <ReactMarkdown components={components}>{props.body}</ReactMarkdown>
      <Byline author={props.author.handle} time={new Date(props.createdAt)} />
    </Flex>
  );
}