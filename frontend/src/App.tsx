import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import { Login } from './components/domain/login/Login';
import { Home } from './components/domain/home/Home';
import { NotFound } from './components/domain/not-found/NotFound';
import { Route, Routes, BrowserRouter } from 'react-router-dom';

function App() {
  const { isAuthenticated, loginWithPopup } = useAuth0();

  return (
    <div className="App">
      <header className="App-header">
        {!isAuthenticated ? (
          <Login loginWithPopup={loginWithPopup} />
        ) : (
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="*" element={<NotFound />} />
            </Routes>
          </BrowserRouter>
        )}
      </header>
    </div>
  );
}

export default App;
