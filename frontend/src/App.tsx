import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import { Login } from './components/domain/login/Login';
import { Home } from './components/domain/home/Home';
import { NotFound } from './components/domain/not-found/NotFound';
import { Route, Routes, BrowserRouter } from 'react-router-dom';
import { SideBar } from './components/base/side-bar/SideBar';
import { Flex } from '@mantine/core';

function App() {
  const { isAuthenticated } = useAuth0();

  return !isAuthenticated ? (
    <div className="App">
      <header className="App-header">
        <Login />
      </header>
    </div>
  ) : (
    <div className="HasSession">
      <Flex
        gap="md"
        justify="flex-start"
        align="center"
        direction="row"
        wrap="wrap"
      >
        <SideBar />
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </Flex>
    </div>
  );
}

export default App;
