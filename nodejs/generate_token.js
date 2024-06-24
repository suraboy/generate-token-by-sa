const { GoogleAuth } = require('google-auth-library');
const path = require('path');

async function generateToken() {
    try {
        const keyFilePath = path.join(__dirname, './storage/paotang-core-sit-d50437475f5b.json'); // Update this path
        const targetAudience = 'https://paotang-telco-check-6ti7spcsya-as.a.run.app'; // Update this URL

        const auth = new GoogleAuth({
            keyFilename: keyFilePath
        });

        const client = await auth.getIdTokenClient(targetAudience);
        const tokenResponse = await client.getRequestHeaders();
        const token = tokenResponse.Authorization.split(' ')[1];

        console.log('ID Token:', token);
    } catch (err) {
        console.error('Error generating token:', err.message);
    }
}

generateToken();
