import { Container } from "@chakra-ui/react";
import { AllGroups } from "./pages/AllGroups";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { GroupInfo } from "./pages/GroupInfo";
import NotFound from './notFound';
function App() {
  return (
    <BrowserRouter>
      <Container>
        <Routes>
          <Route path="/" element={<AllGroups />} />
          <Route path="/:id" element={<GroupInfo />} />
          <Route path='*' element={<NotFound/>}/>
        </Routes>
      </Container>
    </BrowserRouter>
  );
}

export default App;
