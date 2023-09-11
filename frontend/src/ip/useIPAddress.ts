import { useEffect, useMemo, useState } from 'react';
import { IP } from '../env/ip/whiteList';

const FETCH_IP_ADDRESS_URL = 'https://api.ipify.org?format=json';

export const useIPAddress = () => {
  const [ipAddress, setIPAddress] = useState<string>('');

  useEffect(() => {
    fetch(FETCH_IP_ADDRESS_URL)
      .then((response) => response.json())
      .then((data) => {
        setIPAddress(data.ip);
      })
      .catch((error) => console.log(error));
  }, []);

  return { ipAddress };
};

export const useIPAccess = () => {
  const { ipAddress } = useIPAddress();

  const whiteList = new IP().getWhiteList();

  const isValidAccess = useMemo(
    () => whiteList.includes(ipAddress),
    [ipAddress],
  );

  return {
    isValidAccess,
  };
};
