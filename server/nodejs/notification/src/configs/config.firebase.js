
// Import Firebase Admin SDK
const firebase_admin = require('firebase-admin');

// Initialize the app with a service account, granting admin privileges
const serviceAccount = require('../key/music-firebase-key.json'); // Ensure the path is correct

const initFirebaseApp = () => {
    firebase_admin.initializeApp({
        credential: firebase_admin.credential.cert(serviceAccount),
    });
}





module.exports = {
    initFirebaseApp
}
