import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import { Auth0Provider } from './provider/Auth0provider';
import { QueryClientProvider } from './provider/QueryClientProvider';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080';

ReactDOM.render(
  <Auth0Provider>
    <QueryClientProvider>
      <React.StrictMode>
        <App />
      </React.StrictMode>
    </QueryClientProvider>
  </Auth0Provider>,
  document.getElementById('root'),
);
