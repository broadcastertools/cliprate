/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Clip } from '../models/Clip';
import type { SiteConfig } from '../models/SiteConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import { request as __request } from '../core/request';

export class ConfService {

    /**
     * Get site configuration.
     * @returns SiteConfig OK
     * @throws ApiError
     */
    public static getSiteConfiguration(): CancelablePromise<SiteConfig> {
        return __request({
            method: 'GET',
            path: `/siteconfig`,
        });
    }

    /**
     * Get a list of clips.
     * @returns Clip OK
     * @throws ApiError
     */
    public static getClips(): CancelablePromise<Clip> {
        return __request({
            method: 'GET',
            path: `/clips`,
        });
    }

}