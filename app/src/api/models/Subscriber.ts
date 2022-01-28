/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Subscriber = {
    /**
     * Our subscriber identifier.
     */
    subscriber_id: string;
    /**
     * Email address of the subscriber, this is where notifications would be sent too.
     */
    email: string;
    /**
     * The display name of the subscriber.
     */
    display_name: string;
    /**
     * Twitch's identifier for the subscriber.
     */
    twitch_id: string;
    /**
     * A timestamp for when the subscriber logged into this application first.
     */
    joined: string;
    /**
     * Is the subscriber currently subscriber.
     */
    is_subscribed: boolean;
    /**
     * Is the subscriber a mod or owner of the broadcaster.
     */
    is_admin: boolean;
}
