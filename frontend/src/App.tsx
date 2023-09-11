import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import { Login } from './components/domain/login/Login';
import { Home } from './components/domain/home/Home';
import { NotFound } from './components/domain/not-found/NotFound';
import { Route, Routes, BrowserRouter } from 'react-router-dom';
import { SideBar } from './components/base/side-bar/SideBar';
import { Flex } from '@mantine/core';
import { Students } from './components/domain/general/students';
import { useIPAccess } from './ip/useIPAddress';
import { InvalidAccess } from './components/domain/invalid-access/InvalidAccess';
import { memo, useMemo } from 'react';

function App() {
  const { isAuthenticated } = useAuth0();
  const { isValidAccess } = useIPAccess();

  const { components } = useCalculateComponents({
    isAuthenticated,
    isValidAccess,
  });

  return components;
}

export default App;

/**
 * @description
 * IPアドレスによるアクセス制限を行った際に認証されない場合に表示するコンポーネント
 */
const InvalidAccessComponent = memo(() => {
  return (
    <div className="App">
      <header className="App-header">
        <InvalidAccess />
      </header>
    </div>
  );
});

/**
 * @description
 * Auth0による認証されていない場合に表示するコンポーネント
 */
const UnauthenticatedComponent = memo(() => {
  return (
    <div className="App">
      <header className="App-header">
        <Login />
      </header>
    </div>
  );
});

/**
 * @description
 * IP制限、Auth0の認証が共に成功した場合に表示するコンポーネント
 */
const Component = memo(() => {
  return (
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
            <Route path="students" element={<Students />} />
          </Routes>
        </BrowserRouter>
      </Flex>
    </div>
  );
});

const useCalculateComponents = (props: {
  isAuthenticated: boolean;
  isValidAccess: boolean;
}) => {
  const { isAuthenticated, isValidAccess } = props;
  const components = useMemo(
    () =>
      !isValidAccess ? (
        <InvalidAccessComponent />
      ) : !isAuthenticated ? (
        <UnauthenticatedComponent />
      ) : (
        <Component />
      ),
    [isValidAccess, isAuthenticated],
  );

  return { components };
};
