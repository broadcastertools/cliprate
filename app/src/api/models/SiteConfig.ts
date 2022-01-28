/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type SiteConfig = {
    /**
     * The name of the streamer that is hosting this insance, for example MuTeX
     */
    streamer_display_name: string;
    /**
     * The name of the streamer that is hosting this insance, for example mutex
     */
    streamer_login: string;
    /**
     * The name of the streamer that is hosting this insance, for example 98506045
     */
    streamer_id: string;
    /**
     * The expected domain for this application, for example clips.mutexisthegoat.com
     */
    domain: string;
    /**
     * A full URI to a logo that should be displayed in the app.
     */
    logo_uri: string;
    /**
     * A hex color used for the appbar.
     */
    appbar_color: string;
    /**
     * A URL to redirect the user to login.
     */
    authorization_url: string;
    /**
     * When true, subscribers that was gifted their subscription can login.
     */
    is_gifted_authorized: boolean;
}
