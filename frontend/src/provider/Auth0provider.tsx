import { createContext, memo, useContext } from 'react';
import { Auth0Provider as Provider, useAuth0 } from '@auth0/auth0-react';
import {
  DOMAIN_DEV,
  DOMAIN_PROD,
  CLIENT_ID_DEV,
  CLIENT_ID_PROD,
} from '../env/auth0/auth0.env';

interface Props {
  children: JSX.Element;
}

export type Auth0Context = ReturnType<typeof useAuth0>;
const Context = createContext<Auth0Context | undefined>(undefined);
export const Auth0Context = Context;

export const Auth0Provider = memo((props: Props) => {
  const DOMAIN = DOMAIN_DEV || DOMAIN_PROD;
  const CLIENT_ID = CLIENT_ID_DEV || CLIENT_ID_PROD;

  const params = useAuth0();

  return (
    <Provider
      domain={DOMAIN}
      clientId={CLIENT_ID}
      authorizationParams={{ redirect_uri: window.location.origin }}
    >
      <Context.Provider value={params}>{props.children}</Context.Provider>;
    </Provider>
  );
});

export const useAuth = () => {
  return useContext(Context);
};
