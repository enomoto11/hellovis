import { useAuth0 } from '@auth0/auth0-react';
import { memo } from 'react';
import { Button } from '@mantine/core';

export const Home = memo(() => {
  const { logout } = useAuth0();

  return (
    <div>
      <h1>Home</h1>
      <Button
        onClick={() => {
          logout();
        }}
      >
        Log out
      </Button>
    </div>
  );
});
