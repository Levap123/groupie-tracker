import axios from "axios";
import { useEffect, useState } from "react";
import { Box, Heading, Text, Avatar } from "@chakra-ui/react";
import { GroupModel } from "../types";
import { Link } from "react-router-dom";

export const AllGroups = () => {
  const [groups, setGroups] = useState<GroupModel[]>();

  useEffect(() => {
    axios
      .get<GroupModel[]>("http://localhost:8080/api/artists/", {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "*",
          "Access-Control-Allow-Headers": "*"
        }
      })
      .then((response) => setGroups(response.data));
  }, []);
  if (!groups) {
    return <Box padding="50px 10px">...Loading</Box>;
  }

  return (
    <Box padding="20px 10px">
      {groups.map((group) => (
        <Link to={"/" + group.id}>
          <Box
            key={group.id}
            display="flex"
            padding="10px"
            margin="15px 0"
            gap="10px"
            alignItems="center"
            borderRadius="10px"
            boxShadow="md"
          >
            <Avatar size="xl" src={group.image} />
            <Box>
              <Heading>{group.name}</Heading>
              <Text>Created: {group.creationDate}</Text>
            </Box>
          </Box>
        </Link>
      ))}
    </Box>
  );
};
