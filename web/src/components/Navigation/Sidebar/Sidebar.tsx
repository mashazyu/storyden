import { Box, Divider, Flex, VStack } from "@chakra-ui/react";
import { SIDEBAR_WIDTH, useNavigation } from "../useNavigation";
import { Authbar } from "./components/Authbar";
import { CategoryList } from "./components/CategoryList";
import { Title } from "./components/Title";
import { Toolbar } from "./components/Toolbar";

export function Sidebar() {
  const { title } = useNavigation();

  return (
    <Flex
      id="desktop-nav-container"
      as="header"
      position="fixed"
      justifyContent="end"
      width={SIDEBAR_WIDTH}
      height="full"
      bgColor="blackAlpha.50"
    >
      <Box
        id="desktop-nav-box"
        maxWidth="2xs"
        minWidth={{
          base: "full",
          lg: "3xs",
        }}
        height="full"
        p={2}
      >
        <VStack
          as="nav"
          height="full"
          gap={2}
          justifyContent="space-between"
          alignItems="start"
        >
          <VStack width="full" alignItems="start" overflow="hidden">
            <Title>{title}</Title>

            <Toolbar />

            <Divider />

            <Box overflowY="scroll" width="full">
              <CategoryList />
            </Box>
          </VStack>

          <Divider />

          <VStack alignItems="start">
            <Authbar />
          </VStack>
        </VStack>
      </Box>
    </Flex>
  );
}