import { useAuth0 } from '@auth0/auth0-react';
import { memo } from 'react';

export const Home = memo(() => {
  const { logout } = useAuth0();

  return (
    <div>
      <h1>Home</h1>
      <button
        onClick={() => {
          logout();
        }}
      >
        Log out
      </button>
    </div>
  );
});
