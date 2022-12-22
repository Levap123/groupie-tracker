import { useState, useEffect } from "react";
import { Box, Heading, Text, Avatar, Button } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { Concerts, Group, Err } from "../types";
import { Link } from "react-router-dom";

export const GroupInfo = () => {
  const [groupInfo, setGroupInfo] = useState<Group>();
  const [concerts, setConcerts] = useState<Concerts>();
  const [err, setError] = useState<Err>();
  const { id } = useParams();
  function capitalizeFirstLetter(str: string) {
    return str.replace(/(^\w{1})|(\s+\w{1})/g, letter => letter.toUpperCase());
  }
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
      .then((response) => setConcerts(response.data))
      .catch((error) => setError(error.response.data));
  }, [id]);
  if (err) {
    return <Box padding="50px 10px">{err.msg}</Box>;
  }
  if (!groupInfo) {
    return <Box padding="50px 10px">...Loading</Box>;
  }
  if (!concerts) {
    return <Box padding="50px 10px">...Loading</Box>;
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
          verticalAlign="center"
          alignSelf="center"
          borderRadius="10px"
          boxShadow="md"
          width="200%"
        >
          <Avatar size="3xl" src={groupInfo.image} />
          <Box>
            <Heading>{groupInfo.name}</Heading>
            <Text>Created: {groupInfo.creationDate}</Text>
            <Text>First album: {groupInfo.firstAlbum}</Text>
            <Box>
              <br></br>
              <Heading size="md">Members</Heading>
              {groupInfo.members.map((member) => (
                <Text key={member}>{member}</Text>
              ))}
            </Box>
            <Box>
              <br></br>
              <Heading size="md">Concerts</Heading>
              {Object.keys(concerts.datesLocations).map((location) => (
                <Text key={location}>{capitalizeFirstLetter(location.replaceAll("_", " ").replaceAll("-", ", "))}: {concerts.datesLocations[location].join(", ")}</Text>
              ))}
            </Box>
          </Box>
        </Box>
      </Box>
  );
};
