with Ada.Text_IO; use Ada.Text_IO;
with Ada.Numerics.Float_Random; use Ada.Numerics.Float_Random;
with Random_Seeds; use Random_Seeds;
with Ada.Real_Time; use Ada.Real_Time;
with Ada.Calendar;

procedure Travelers3 is
  -- Travelers moving on the board
  Nr_Of_Travelers : constant Integer := 15;

  Min_Steps : constant Integer := 10;
  Max_Steps : constant Integer := 100;

  Min_Delay : constant Duration := 0.01;
  Max_Delay : constant Duration := 0.05;
  Move_Timeout : constant Duration := Max_Delay * 10.0; -- Timeout for deadlock detection

  -- 2D Board with torus topology
  Board_Width  : constant Integer := 15;
  Board_Height : constant Integer := 15;

  -- Timing
  Start_Time : Ada.Real_Time.Time := Ada.Real_Time.Clock;  -- global starting time

  -- Random seeds for the tasks' random number generators
  Seeds : Seed_Array_Type(1..Nr_Of_Travelers) := Make_Seeds(Nr_Of_Travelers);

  -- Types, procedures and functions
  -- Positions on the board
  type Position_Type is record	
    X: Integer range 0 .. Board_Width - 1; 
    Y: Integer range 0 .. Board_Height - 1; 
  end record;
  
  -- Direction type for consistent movement
  type Direction_Type is (Up, Down, Left, Right);
  
  -- Protected object for field coordination
  protected type Field_Type is
    entry Request_Entry;
    procedure Release;
    function Is_Occupied return Boolean;
  private
    Occupied : Boolean := False;
  end Field_Type;
  
  protected body Field_Type is
    entry Request_Entry when not Occupied is
    begin
      Occupied := True;
    end Request_Entry;
    
    procedure Release is
    begin
      Occupied := False;
    end Release;
    
    function Is_Occupied return Boolean is
    begin
      return Occupied;
    end Is_Occupied;
  end Field_Type;
  
  -- Board of fields
  type Board_Type is array(0 .. Board_Width - 1, 0 .. Board_Height - 1) of Field_Type;
  Board : Board_Type;

  -- elementary steps
  procedure Move_Down(Position: in out Position_Type) is
  begin
    Position.Y := (Position.Y + 1) mod Board_Height;
  end Move_Down;

  procedure Move_Up(Position: in out Position_Type) is
  begin
    Position.Y := (Position.Y + Board_Height - 1) mod Board_Height;
  end Move_Up;

  procedure Move_Right(Position: in out Position_Type) is
  begin
    Position.X := (Position.X + 1) mod Board_Width;
  end Move_Right;

  procedure Move_Left(Position: in out Position_Type) is
  begin
    Position.X := (Position.X + Board_Width - 1) mod Board_Width;
  end Move_Left;

  -- traces of travelers
  type Trace_Type is record 	      
    Time_Stamp : Duration;	      
    Id : Integer;
    Position : Position_Type;      
    Symbol : Character;	      
  end record;	      

  type Trace_Array_type is array(0 .. Max_Steps) of Trace_Type;

  type Traces_Sequence_Type is record
    Last : Integer := -1;
    Trace_Array : Trace_Array_type;
  end record; 

  procedure Print_Trace(Trace : Trace_Type) is
    Symbol_Str : String(1..1);
  begin
    Symbol_Str(1) := Trace.Symbol;
    Put_Line(
        Duration'Image(Trace.Time_Stamp) & " " &
        Integer'Image(Trace.Id) & " " &
        Integer'Image(Trace.Position.X) & " " &
        Integer'Image(Trace.Position.Y) & " " &
        Symbol_Str
      );
  end Print_Trace;

  procedure Print_Traces(Traces : Traces_Sequence_Type) is
  begin
    for I in 0 .. Traces.Last loop
      Print_Trace(Traces.Trace_Array(I));
    end loop;
  end Print_Traces;

  -- task Printer collects and prints reports of traces
  task Printer is
    entry Report(Traces : Traces_Sequence_Type);
    entry Done;
  end Printer;
  
  task body Printer is 
    Travelers_Reported : Integer := 0;
  begin
    loop
      select
        accept Report(Traces : Traces_Sequence_Type) do
          Print_Traces(Traces);
          Travelers_Reported := Travelers_Reported + 1;
        end Report;
      or
        accept Done;
        exit when Travelers_Reported = Nr_Of_Travelers;
      or
        terminate;
      end select;
    end loop;
  end Printer;

  -- travelers
  type Traveler_Type is record
    Id : Integer;
    Symbol : Character;
    Position : Position_Type;
    Direction : Direction_Type;    
  end record;

  task type Traveler_Task_Type is	
    entry Init(Id: Integer; Seed: Integer; Symbol: Character);
    entry Start;
  end Traveler_Task_Type;	

  task body Traveler_Task_Type is
    G : Generator;
    Traveler : Traveler_Type;
    Time_Stamp : Duration;
    Nr_of_Steps : Integer;
    Traces : Traces_Sequence_Type;
    Deadlocked : Boolean := False;

    procedure Store_Trace is
    begin  
      Traces.Last := Traces.Last + 1;
      Traces.Trace_Array(Traces.Last) := ( 
          Time_Stamp => Time_Stamp,
          Id => Traveler.Id,
          Position => Traveler.Position,
          Symbol => Traveler.Symbol
        );
    end Store_Trace;
    
    -- Attempt to move in the predetermined direction with deadlock detection
    function Try_Move return Boolean is
      New_Position : Position_Type;
      Start_Wait_Time : Ada.Calendar.Time;
      Elapsed : Duration;
    begin
      New_Position := Traveler.Position;
      
      -- Move in the pre-determined direction
      case Traveler.Direction is
        when Up =>
          Move_Up(New_Position);
        when Down =>
          Move_Down(New_Position);
        when Left =>
          Move_Left(New_Position);
        when Right => 
          Move_Right(New_Position);
      end case;
      
      -- Try to acquire the field with timeout
      Start_Wait_Time := Ada.Calendar.Clock;
      
      select
        Board(New_Position.X, New_Position.Y).Request_Entry;
        -- Field acquired, release the old one
        Board(Traveler.Position.X, Traveler.Position.Y).Release;
        Traveler.Position := New_Position;
        return True;
      or
        delay Move_Timeout;
        -- Deadlock suspected
        Elapsed := Ada.Calendar."-"(Ada.Calendar.Clock, Start_Wait_Time);
        Put_Line("Traveler " & Integer'Image(Traveler.Id) & 
                 " potential deadlock after " & Duration'Image(Elapsed) & " seconds");
        -- Change to lowercase symbol to indicate deadlock
        Traveler.Symbol := Character'Val(Character'Pos(Traveler.Symbol) + 32);
        Deadlocked := True;
        return False;
      end select;
    end Try_Move;

  begin
    accept Init(Id: Integer; Seed: Integer; Symbol: Character) do
      Reset(G, Seed); 
      Traveler.Id := Id;
      Traveler.Symbol := Symbol;
      
      -- Set initial position to (i, i) on the diagonal
      Traveler.Position := (X => Id, Y => Id);
      
      -- Request and occupy the initial position
      Board(Traveler.Position.X, Traveler.Position.Y).Request_Entry;
      
      -- Determine direction based on ID parity
      -- Even IDs move vertically (Up or Down)
      -- Odd IDs move horizontally (Left or Right)
      if Id mod 2 = 0 then
        -- Even ID: vertical movement (up or down)
        if Random(G) < 0.5 then
          Traveler.Direction := Up;
        else
          Traveler.Direction := Down;
        end if;
      else
        -- Odd ID: horizontal movement (left or right)
        if Random(G) < 0.5 then
          Traveler.Direction := Left;
        else
          Traveler.Direction := Right;
        end if;
      end if;
      
      Time_Stamp := To_Duration(Ada.Real_Time.Clock - Start_Time);
      Store_Trace; -- store starting position
      
      -- Number of steps to be made by the traveler  
      Nr_of_Steps := Min_Steps + Integer(Float(Max_Steps - Min_Steps) * Random(G));
    end Init;
    
    accept Start do
      null;
    end Start;

    for Step in 1 .. Nr_of_Steps loop
      exit when Deadlocked;
      
      delay Min_Delay + (Max_Delay - Min_Delay) * Duration(Random(G));
      
      if not Try_Move then
        -- Deadlock detected, update trace with lowercase symbol
        Time_Stamp := To_Duration(Ada.Real_Time.Clock - Start_Time);
        Store_Trace;
        exit;
      end if;
      
      Time_Stamp := To_Duration(Ada.Real_Time.Clock - Start_Time);
      Store_Trace;
    end loop;
    
    -- Report traces and release the final position
    Printer.Report(Traces);
    
    if not Deadlocked then
      -- Normal termination, release final position
      Board(Traveler.Position.X, Traveler.Position.Y).Release;
    end if;
    
    Printer.Done;
  end Traveler_Task_Type;

  Travel_Tasks : array(0 .. Nr_Of_Travelers - 1) of Traveler_Task_Type;
  Symbol : Character := 'A';
begin 
  -- Print the line with the parameters needed for display script:
  Put_Line(
      "-1 " &
      Integer'Image(Nr_Of_Travelers) & " " &
      Integer'Image(Board_Width) & " " &
      Integer'Image(Board_Height)      
    );

  -- Initialize traveler tasks
  for I in Travel_Tasks'Range loop
    Travel_Tasks(I).Init(I, Seeds(I+1), Symbol);
    Symbol := Character'Succ(Symbol);
  end loop;

  -- Start traveler tasks
  for I in Travel_Tasks'Range loop
    Travel_Tasks(I).Start;
  end loop;
end Travelers3;