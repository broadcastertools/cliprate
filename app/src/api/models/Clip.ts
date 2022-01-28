/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Clip = {
    /**
     * Our subscriber identifier.
     */
    clip_id: string;
    /**
     * Identifier of the subscriber.
     */
    posted_by: string;
    /**
     * Timestamp for when the clip was posted.
     */
    posted: string;
    /**
     * A display title.
     */
    title: string;
    /**
     * The category name the clip is posted in.
     */
    category: string;
    /**
     * A URI to the thumbnail.
     */
    thumbnail_uri: string;
    /**
     * Where the clip is hosted.
     */
    type: Clip.type;
    /**
     * Data for the clip, this should be used along with clip type.
     */
    data: string;
}

export namespace Clip {

    /**
     * Where the clip is hosted.
     */
    export enum type {
        YOUTUBE = 'youtube',
        TWITCHCLIP = 'twitchclip',
    }


}
