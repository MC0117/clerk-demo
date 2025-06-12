import { Routes, Route, Link } from 'react-router-dom'
import { SignedIn, SignedOut, SignInButton, UserButton } from '@clerk/clerk-react'
import type { ReactNode } from 'react'
import './styles/App.css'

function PublicPage() {
  return (
    <div className="page">
      <h1>Welcome to Our App</h1>
      <p>This is a public page that anyone can see.</p>
      <div style={{ marginTop: '2rem' }}>
        <SignedOut>
          <div>
            <p>Sign in to access your profile!</p>
            <SignInButton mode="modal" />
          </div>
        </SignedOut>
        <SignedIn>  
          <p>You're signed in! Check out your <Link to="/profile">profile</Link>.</p>
        </SignedIn>
      </div>
    </div>
  )
}

function ProfilePage() {
  return (
    <div className="page">
      <h1>Your Profile</h1>
      <SignedIn>
        <div>
          <p>Welcome to your profile!</p>
          <div style={{ marginTop: '1rem' }}>
            <UserButton afterSignOutUrl="/" />
          </div>
        </div>
      </SignedIn>
    </div>
  )
} 

function Navigation() {
  return (
    <nav>
      <div className="nav-left">
        <Link to="/">Home</Link>
        <SignedIn>
          <Link to="/profile">Profile</Link>
        </SignedIn>
      </div>
      <div className="nav-right">
        <SignedOut>
          <SignInButton mode="modal" />
        </SignedOut>
        <SignedIn>
          <UserButton afterSignOutUrl="/" />
        </SignedIn>
      </div>
    </nav>
  )
}

function ProtectedRoute({ children }: { children: ReactNode }) {
  return (
    <>
      <SignedIn>
        {children}
      </SignedIn>
      <SignedOut>
        <div className="page">
          <h1>Please Sign In</h1>
          <SignInButton mode="modal" />
        </div>
      </SignedOut>
    </>
  )
}

function App() {
  return (
    <div className="app">
      <Navigation />
      <main>
        <Routes>
          <Route path="/" element={<PublicPage />} />
          <Route 
            path="/profile" 
            element={
              <ProtectedRoute>
                <ProfilePage />
              </ProtectedRoute>
            } 
          />
        </Routes>
      </main>
    </div>
  )
}

export default App
