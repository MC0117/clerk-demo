import express from 'express';
import { ClerkExpressRequireAuth } from '@clerk/clerk-sdk-node';
import cors from 'cors';

const app = express();
app.use(cors());
app.use(express.json());

// Public endpoint - no auth required
app.get('/api/public', (req, res) => {
  res.json({ message: 'This is public data' });
});

// Protected endpoint - requires authentication
app.get('/api/protected', ClerkExpressRequireAuth(), (req, res) => {
  // req.auth.userId contains the authenticated user's ID
  res.json({ 
    message: 'This is protected data',
    userId: req.auth.userId
  });
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
}); 