import { createContext, useContext, useEffect, useState } from "react";
import { SiteConfig } from "../api/models/SiteConfig";
import { Alert } from "@mui/material";

export const ConfigContext = createContext<SiteConfig|null>(null);


export const ConfigProvider: React.FC<{}> = ({
    children,
}) => {
  const [err, setErr] = useState<string|null>(null);
  const [config, setConfig] = useState<SiteConfig|null>(null);

  useEffect(() => {
      try {
          fetch('/v1/siteconfig')
              .then(async response => {
                  setConfig(await response.json());
              })
              .catch(e => {
                  console.error(e);
                  setErr("There was an issue loading the configuration, please try again later.");
              });
          
      } catch (e) {
        console.error(e);
        setErr("Please check your network connection.");
      }
  }, [setConfig]);

  if (err != null) {
    return <Alert severity="error">{err}</Alert>
  }

  if (config === null) {
    return null;
  }

  return (
      <ConfigContext.Provider value={config}>
      {children}
      </ConfigContext.Provider>
  );
  };

export function useConfigContext() {
    const ctx = useContext(ConfigContext);
  
    if (!ctx) {
      throw new Error("Using ConfigContext outside of the provider.");
    }
  
    return ctx;
  }
