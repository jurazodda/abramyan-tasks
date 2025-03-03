import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        double p = 3.14;
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter the diameter of the circle: ");
        double d = scanner.nextDouble();

        double l = p * d;
        System.out.printf("Circumference of the circle: %.2f\n", l);
    }
}