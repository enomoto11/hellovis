import React from 'react';
import './index.css';
import App from './App';
import { Auth0Provider } from './provider/Auth0provider';
import { QueryClientProvider } from './provider/QueryClientProvider';
import axios from 'axios';
import { createRoot } from 'react-dom/client';

axios.defaults.baseURL = 'http://localhost:8080';

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
const root = createRoot(document.getElementById('root')!);
root.render(
  <React.StrictMode>
    <Auth0Provider>
      <QueryClientProvider>
        <App />
      </QueryClientProvider>
    </Auth0Provider>
  </React.StrictMode>,
);
