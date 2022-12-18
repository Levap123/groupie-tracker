import { useState, useEffect } from "react";
import { Box, Heading, Text, Avatar, Button } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { Concerts, Group } from "../types";
import { Link } from "react-router-dom";

export const GroupInfo = () => {
  const [groupInfo, setGroupInfo] = useState<Group>();
  const [concerts, setConcerts] = useState<Concerts>();
  const { id } = useParams();

  useEffect(() => {
    axios
      .get<Group>("http://localhost:8080/api/artists/" + id, {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "*",
          "Access-Control-Allow-Headers": "*"
        }
      })
      .then((response) => setGroupInfo(response.data));

    axios
      .get<Concerts>("http://localhost:8080/api/concerts/" + id, {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "*",
          "Access-Control-Allow-Headers": "*"
        }
      })
      .then((response) => setConcerts(response.data));
  }, [id]);
  console.log(groupInfo)
  if (!groupInfo || !concerts) {
    return <Box padding="50px 10px">...Loading</Box>;
  }

  function replace(location: string): import("react").ReactNode {
    throw new Error("Function not implemented.");
  }

  return (
    <Box padding="20px 10px">
      <Link to="/">
        <Button width="100%">Back</Button>
      </Link>
      <Box
        key={groupInfo.id}
        display="flex"
        padding="10px"
        margin="15px 0"
        gap="10px"
        alignItems="center"
        borderRadius="10px"
        boxShadow="md"
      >
        <Avatar size="2xl" src={groupInfo.image} />
        <Box>
          <Heading>{groupInfo.name}</Heading>
          <Text>Created: {groupInfo.creationDate}</Text>
          <Text>First album: {groupInfo.firstAlbum}</Text>
          <Box>
            <Heading size="md">Members</Heading>
            {groupInfo.members.map((member) => (
              <Text key={member}>{member}</Text>
            ))}
          </Box>
          <Box>
            <Heading size="md">Concerts</Heading>
            {Object.keys(concerts.datesLocations).map((location) => (
              <Text key={location}>{location.replaceAll("_", " ").replace("-",", ")}: {concerts.datesLocations[location]}</Text>
            ))}
          </Box>
        </Box>
      </Box>
    </Box>
  );
};
