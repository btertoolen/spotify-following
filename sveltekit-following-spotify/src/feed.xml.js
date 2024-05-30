import songs from './_songs.js'

export async function get() {
    return {
        body: songs
    };
}