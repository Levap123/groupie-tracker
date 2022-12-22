import { Box } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

const NotFound = () => (
    <Box> <h1>404 - Not Found!</h1>
        <Link to="/">Go Home</Link></Box>
);

export default NotFound;