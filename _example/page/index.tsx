
type IndexProps = {
    name: string;
};

const IndexPage = (p: IndexProps) => {
  return (
    <div>
      <h1>Welcome to the Index Page</h1>
      <p>Hello, {p.name}!</p>
    </div>
  );
};

export default IndexPage;