import { memo } from 'react';
import { Auth0Provider as Provider } from '@auth0/auth0-react';
import {
  DOMAIN_DEV,
  DOMAIN_PROD,
  CLIENT_ID_DEV,
  CLIENT_ID_PROD,
} from '../env/auth0.env';

interface Props {
  children: JSX.Element;
}

export const Auth0Provider = memo((props: Props) => {
  const DOMAIN = DOMAIN_DEV || DOMAIN_PROD;
  const CLIENT_ID = CLIENT_ID_DEV || CLIENT_ID_PROD;

  return (
    <Provider
      domain={DOMAIN}
      clientId={CLIENT_ID}
      authorizationParams={{ redirect_uri: window.location.origin }}
    >
      {props.children || <></>}
    </Provider>
  );
});
