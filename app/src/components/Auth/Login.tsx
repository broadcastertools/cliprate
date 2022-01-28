import React, { useEffect, useState } from "react";
import { Box } from "@mui/system";
import { Button, Card, CardActions, CardContent, CircularProgress, Container, Typography } from "@mui/material";
import { AuthService } from "../../api/services/AuthService";
import { useConfigContext } from "../../siteconfig";

export const LoginScreen = () => {
    const cnf = useConfigContext();
    const logo = cnf.logo_uri;
    const streamerDisplayName = cnf.streamer_display_name;
    const authorizationURL = cnf.authorization_url;
    const allowGifted = cnf.is_gifted_authorized;

    return <Container sx={{mt: 2}} maxWidth="sm">
        <Box textAlign="center">
            <img src={logo} alt={`${streamerDisplayName}'s logo`} />
        </Box>
        <Card sx={{mt: 1}}>
            <CardContent>
                Please login to the clip rate app using your Twitch account.

                <ul>
                    <li><strong>We DO NOT</strong> ever have access to your Twitch login.</li>
                    { allowGifted ?
                        <li><strong>You MUST</strong> be a subscriber and it cannot be gifted.</li> :
                        <li><strong>You CAN</strong> be a subscriber and it be gifted.</li>
                    }
                    <li><strong>You MUST</strong> read all rules when using this app otherwise you'll be banned.</li>
                </ul>
            </CardContent>
            <CardActions>
                <Button
                    variant="contained"
                    href={authorizationURL}
                >Login with Twitch</Button>
            </CardActions>
        </Card>
    </Container>
};

export const VerifyCode = () => {
    const cnf = useConfigContext();
    const logo = cnf.logo_uri;
    const streamerDisplayName = cnf.streamer_display_name;

    const u = new URLSearchParams(window.location.search);
    const code = u.get("code");
    const [message, setMessage] = useState<null|string>(null);

    useEffect(() => {
        if (code === null) {
            console.error("code cannot be null");
            return;
        }

        console.error(`requesting code for ${code}`);
        AuthService.loginWithTwitchCode({code})
            .then(res => {
                console.error('OK: ', `${res}`);
                setMessage(res.token);
            })
            .catch(err => {
                console.error('FAIL: ', `${err}`);
                setMessage(`${JSON.parse(err.body).message}`);
            })
    }, [code]);


    return <Container sx={{mt: 2}} maxWidth="sm">
        <Box textAlign="center">
            <img src={logo} alt={`${streamerDisplayName}'s logo`} />
        </Box>
        <Card sx={{mt: 1}}>
            <CardContent>
                We're just verifying that your login, if this is your first time this may take a bit longer.
            </CardContent>
            <CardActions>
                { message === null ? <CircularProgress /> :
                    <Typography variant="body1">{message}</Typography> }
            </CardActions>
        </Card>
    </Container>
};
