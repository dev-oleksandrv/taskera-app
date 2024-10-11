import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function RegisterPage() {
  return (
      <form className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="name">Username</Label>
          <Input
            id="username"
            placeholder="Enter your name"
            required
            name="username"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            placeholder="Enter your email"
            required
            type="email"
            name="email"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            required
            type="password"
            name="password"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="confirm-password">Confirm Password</Label>
          <Input
            id="confirmPassword"
            required
            type="password"
            name="confirmPassword"
            className="border-teal-200 focus:border-teal-500"
          />
        </div>
        <Button
          className="w-full bg-teal-600 text-white hover:bg-teal-700"
          type="submit"
        >
          Register
        </Button>
      </form>
  );
}
