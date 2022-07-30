import { useEffect, useState } from 'react';

const ENV = process.env.REACT_APP_ENV || '';
const API_URL = ENV === 'dev' ? 'http://localhost:18080' : '';

const App = () => {
  const [name, setName] = useState('Go');
  const [message, setMessage] = useState('');

  useEffect(() => {
    (async () => {
      try {
        const rsp = await fetch(`${API_URL}/hello/${encodeURIComponent(name)}`);
        if (!rsp.ok) throw new Error(`${rsp.status} ${rsp.statusText}`);
        const data = await rsp.json();
        if (data.message) setMessage(data.message);
      } catch (error) {
        if (error instanceof Error) console.error(error.message);
        setMessage('');
      }
    })();
  }, [name]);

  return (
    <div>
      <div>
        <input
          value={name}
          onChange={e => setName(e.target.value)}
        />
      </div>
      <div>
        {message}
      </div>
    </div>
  );
};

export default App;
