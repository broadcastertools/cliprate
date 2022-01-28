/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { loginWithTwitchCodeRequest } from '../models/loginWithTwitchCodeRequest';
import type { loginWithTwitchCodeResponse } from '../models/loginWithTwitchCodeResponse';
import type { Subscriber } from '../models/Subscriber';
import type { CancelablePromise } from '../core/CancelablePromise';
import { request as __request } from '../core/request';

export class AuthService {

    /**
     * Get current logged in subscriber.
     * @returns Subscriber OK
     * @throws ApiError
     */
    public static getSelf(): CancelablePromise<Subscriber> {
        return __request({
            method: 'GET',
            path: `/me`,
        });
    }

    /**
     * Login with Twitch code.
     * @param requestBody
     * @returns loginWithTwitchCodeResponse OK
     * @throws ApiError
     */
    public static loginWithTwitchCode(
        requestBody: loginWithTwitchCodeRequest,
    ): CancelablePromise<loginWithTwitchCodeResponse> {
        return __request({
            method: 'POST',
            path: `/login`,
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `OK`,
            },
        });
    }

}