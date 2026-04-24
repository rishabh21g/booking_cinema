
export const Header = ({ userID }: { userID: string }) => (
  <header>
    <h1>Cinema Booking</h1>
    <div className="user-id">user: {userID}</div>
  </header>
);